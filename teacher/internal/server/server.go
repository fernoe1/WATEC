package server

import (
	"context"
	"log/slog"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/fernoe1/WATEC/teacher/config"
	teacherSvc "github.com/fernoe1/WATEC/teacher/internal/application/adapter/grpc"
	"github.com/fernoe1/WATEC/teacher/internal/application/repository"
	"github.com/fernoe1/WATEC/teacher/internal/application/usecase"
	tchersvc "github.com/fernoe1/protogen/watec/service/teacher"
	"github.com/redis/go-redis/v9"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"gorm.io/gorm"
)

type Server struct {
	gormDB *gorm.DB
	redis  *redis.Client
	cfg    *config.Config
}

func NewServer(
	gormDB *gorm.DB,
	redis *redis.Client,
	cfg *config.Config,
) *Server {
	return &Server{
		gormDB: gormDB,
		redis:  redis,
		cfg:    cfg,
	}
}

func (s *Server) Run() error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	teacherRepository := repository.NewGormRepository(s.gormDB)
	teacherCache := repository.NewRedisRepository(s.redis)

	teacherUC := usecase.NewTeacherUsecase(teacherRepository, teacherCache)

	grpcAddr := s.cfg.Server.Port

	listener, err := net.Listen("tcp", grpcAddr)
	defer listener.Close()
	if err != nil {
		return err
	}

	grpcServer := grpc.NewServer(
		grpc.KeepaliveParams(keepalive.ServerParameters{
			MaxConnectionIdle: s.cfg.Server.MaxConnectionIdle * time.Minute,
			MaxConnectionAge:  s.cfg.Server.MaxConnectionAge * time.Minute,
			Time:              s.cfg.Server.Timeout * time.Minute,
			Timeout:           s.cfg.Server.Timeout * time.Second,
		}),
	)

	teacherSvc := teacherSvc.NewTeacherService(teacherUC)
	tchersvc.RegisterTeacherServiceServer(grpcServer, teacherSvc)

	go func() {
		if err := grpcServer.Serve(listener); err != nil {
			slog.Error("gRPC server error", "error", err)
			cancel()
		}
		slog.Info("gRPC server listening", "on", grpcAddr)
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	select {
	case sig := <-sigChan:
		slog.Warn("shutting down", "sig", sig)
	case <-ctx.Done():
		slog.Warn("context canceled, shutting down")
	}

	grpcServer.GracefulStop()

	slog.Info("server shutdown")

	return nil
}

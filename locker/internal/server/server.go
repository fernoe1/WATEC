package server

import (
	"context"
	"log/slog"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/fernoe1/WATEC/locker/config"
	lockerGrpc "github.com/fernoe1/WATEC/locker/internal/application/adapter/grpc"
	"github.com/fernoe1/WATEC/locker/internal/application/repository"
	"github.com/fernoe1/WATEC/locker/internal/application/usecase"
	lokrsvc "github.com/fernoe1/protogen/watec/service/locker"
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

	lockerRepository := repository.NewGormRepository(s.gormDB)
	lockerCache := repository.NewRedisRepository(s.redis)

	lockerUC := usecase.NewLockerUsecase(lockerRepository, lockerCache)

	grpcAddr := s.cfg.Server.Port

	listener, err := net.Listen("tcp", grpcAddr)
	defer listener.Close()
	if err != nil {
		return err
	}

	grpcServer := grpc.NewServer(
		grpc.KeepaliveParams(keepalive.ServerParameters{
			MaxConnectionIdle: s.cfg.Server.MaxConnectionIdle * time.Minute,
			Timeout:           s.cfg.Server.Timeout * time.Second,
			MaxConnectionAge:  s.cfg.Server.MaxConnectionAge * time.Minute,
			Time:              s.cfg.Server.Timeout * time.Minute,
		}),
	)

	lockerSvc := lockerGrpc.NewLockerService(lockerUC)
	lokrsvc.RegisterLockerServiceServer(grpcServer, lockerSvc)

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

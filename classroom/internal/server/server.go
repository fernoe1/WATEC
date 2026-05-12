package server

import (
	"context"
	"log/slog"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/fernoe1/WATEC/classroom/config"
	classroomGrpc "github.com/fernoe1/WATEC/classroom/internal/application/adapter/grpc"
	"github.com/fernoe1/WATEC/classroom/internal/application/repository"
	"github.com/fernoe1/WATEC/classroom/internal/application/usecase"
	clsrmsvc "github.com/fernoe1/protogen/watec/service/classroom"
	"github.com/redis/go-redis/v9"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"gorm.io/gorm"
)

type Server struct {
	tracer *trace.Tracer
	log    *slog.Logger
	meter  *metric.Meter
	gormDB *gorm.DB
	redis  *redis.Client
	cfg    *config.Config
}

func NewServer(
	tracer *trace.Tracer,
	log *slog.Logger,
	meter *metric.Meter,
	gormDB *gorm.DB,
	redis *redis.Client,
	cfg *config.Config,
) *Server {
	return &Server{
		tracer: tracer,
		log:    log,
		meter:  meter,
		gormDB: gormDB,
		redis:  redis,
		cfg:    cfg,
	}
}

func (s *Server) Run() error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	classroomRepository := repository.NewGormRepository(s.log, s.gormDB)
	classroomCache := repository.NewRedisRepository(s.log, s.redis)

	classroomUC := usecase.NewClassroomUsecase(s.log, classroomRepository, classroomCache)

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
		grpc.StatsHandler(otelgrpc.NewServerHandler()),
	)

	classroomSvc := classroomGrpc.NewClassroomService(s.log, classroomUC)
	clsrmsvc.RegisterClassroomServiceServer(grpcServer, classroomSvc)

	go func() {
		s.log.Info("gRPC server listening", "on", grpcAddr)
		if err := grpcServer.Serve(listener); err != nil {
			s.log.Error("gRPC server error", "error", err)
			cancel()
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	select {
	case sig := <-sigChan:
		s.log.Warn("shutting down", "sig", sig)
	case <-ctx.Done():
		s.log.Warn("context canceled, shutting down")
	}

	grpcServer.GracefulStop()

	s.log.Info("server shutdown")

	return nil
}

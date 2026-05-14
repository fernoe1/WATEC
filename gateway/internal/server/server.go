package server

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/fernoe1/WATEC/gateway/config"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/trace"
)

type Server struct {
	tracer *trace.Tracer
	log    *slog.Logger
	meter  *metric.Meter
	redis  *redis.Client
	gin    *gin.Engine
	cfg    *config.Config
}

func NewServer(
	tracer *trace.Tracer,
	log *slog.Logger,
	meter *metric.Meter,
	redis *redis.Client,
	cfg *config.Config,
) *Server {
	return &Server{
		tracer: tracer,
		log:    log,
		meter:  meter,
		redis:  redis,
		gin:    gin.New(),
		cfg:    cfg,
	}
}

func (s *Server) Run() error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		if err := s.runHTTPServer(); err != nil {
			slog.Error("http server error", "error", err)
			cancel()
		}
		slog.Info("http server listening", "on", s.cfg.Http.Port)
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	select {
	case sig := <-sigChan:
		s.log.Warn("shutting down", "sig", sig)
	case <-ctx.Done():
		s.log.Warn("context canceled, shutting down")
	}

	if err := s.close(); err != nil {
		slog.Error("http server shutdown error", "error", err)
	}

	slog.Info("server shutdown")

	return nil
}

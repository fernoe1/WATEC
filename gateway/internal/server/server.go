package server

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

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
	srv    *http.Server
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
		slog.Info("http server listening", "on", s.cfg.Http.Port)
		if err := s.runHTTPServer(); err != nil {
			slog.Error("http server error", "error", err)
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

	sCtx, sCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer sCancel()

	if err := s.srv.Shutdown(sCtx); err != nil {
		s.log.Error("http server shutdown error", "error", err)
	}

	s.log.Info("server shutdown")

	return nil
}

package server

import (
	"log/slog"

	"github.com/fernoe1/WATEC/gateway/config"
	"github.com/redis/go-redis/v9"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/trace"
)

type Server struct {
	tracer *trace.Tracer
	log    *slog.Logger
	meter  *metric.Meter
	redis  *redis.Client
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
		cfg:    cfg,
	}
}

func (s *Server) Run() error {
}

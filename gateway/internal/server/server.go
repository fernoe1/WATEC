package server

import (
	"log/slog"

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
}

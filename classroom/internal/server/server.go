package server

import (
	"log/slog"

	"github.com/fernoe1/WATEC/classroom/config"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/trace"
)

type Server struct {
	tracer *trace.Tracer
	log    *slog.Logger
	meter  *metric.Meter
	cfg    *config.Config
}

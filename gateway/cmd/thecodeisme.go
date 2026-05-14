package main

import (
	"context"
	"log"
	"log/slog"

	"github.com/fernoe1/WATEC/gateway/config"
	"github.com/fernoe1/WATEC/gateway/internal/server"
	"github.com/fernoe1/WATEC/gateway/pkg/nats"
	"github.com/fernoe1/WATEC/gateway/pkg/redis"
	"github.com/fernoe1/WATEC/gateway/pkg/telemetry"
	"go.opentelemetry.io/contrib/bridges/otelslog"
	"go.opentelemetry.io/otel"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	slog.Info("reading config")
	cfg, err := config.ReadConfig()
	if err != nil {
		log.Fatal(err)
	}

	slog.Info("initializing otel providers")
	tp, mp, lp, err := telemetry.Init(ctx, cfg)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		tp.Shutdown(ctx)
		mp.Shutdown(ctx)
		lp.Shutdown(ctx)
	}()

	tracer := otel.Tracer(cfg.Telemetry.Name)

	meter := mp.Meter(cfg.Telemetry.Name)

	handler := otelslog.NewHandler(cfg.Telemetry.Name)
	logger := slog.New(handler)
	slog.Info("otel providers set")

	redisClient := redis.NewRedis(cfg)
	slog.Info("redis connected")

	nc := nats.NewNATS(cfg)
	slog.Info("nats connected")

	s := server.NewServer(&tracer, logger, &meter, redisClient, nc, cfg)

	log.Fatal(s.Run())
}

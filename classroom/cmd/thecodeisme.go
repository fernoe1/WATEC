package main

import (
	"context"
	"log"
	"log/slog"

	"github.com/fernoe1/WATEC/classroom/config"
	"github.com/fernoe1/WATEC/classroom/internal/server"
	"github.com/fernoe1/WATEC/classroom/migrate"
	"github.com/fernoe1/WATEC/classroom/pkg/gorm"
	"github.com/fernoe1/WATEC/classroom/pkg/redis"
	"github.com/fernoe1/WATEC/classroom/pkg/telemetry"
	"go.opentelemetry.io/contrib/bridges/otelslog"
	"go.opentelemetry.io/otel"
)

func main() {
	slog.Info("reading config")
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
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

	gormDB := gorm.NewGormDB(cfg)
	slog.Info("gorm connected")
	migrate.Migrate(gormDB)

	redisClient := redis.NewRedis(cfg)
	slog.Info("redis connected")

	s := server.NewServer(&tracer, logger, &meter, gormDB, redisClient, cfg)
	log.Fatal(s.Run())
}

package main

import (
	"context"
	"log"
	"log/slog"

	"github.com/fernoe1/WATEC/classroom/config"
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

	// tracer
	tracer := otel.Tracer("classroom")

	// meter
	meter := mp.Meter("classroom")

	// logger
	handler := otelslog.NewHandler("classroom")
	log := slog.New(handler)
}

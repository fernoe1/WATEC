package main

import (
	"context"
	"log"
	"log/slog"

	"github.com/fernoe1/WATEC/classroom/config"
	"github.com/fernoe1/WATEC/classroom/pkg/telemetry"
	"go.opentelemetry.io/contrib/bridges/otelslog"
)

func main() {
	slog.Info("reading config")
	ctx := context.Background()
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
	tracer := tp.Tracer("classroom")

	// meter
	meter := mp.Meter("classroom")

	// logger
	handler := otelslog.NewHandler("classroom")
	logger := slog.New(handler)
}

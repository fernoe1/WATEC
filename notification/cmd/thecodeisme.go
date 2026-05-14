package main

import (
	"log"
	"log/slog"

	"github.com/fernoe1/WATEC/notification/config"
	"github.com/fernoe1/WATEC/notification/internal/server"
	"github.com/fernoe1/WATEC/notification/pkg/nats"
)

func main() {
	slog.Info("reading config")
	cfg, err := config.ReadConfig()
	if err != nil {
		log.Fatal(err)
	}

	nc := nats.NewNATS()
	slog.Info("nats connected")

	s := server.NewServer(nc, cfg)
	log.Fatal(s.Run())
}

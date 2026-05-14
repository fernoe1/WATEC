package nats

import (
	"log"

	"github.com/fernoe1/WATEC/gateway/config"
	"github.com/nats-io/nats.go"
)

func NewNATS(cfg *config.Config) *nats.Conn {
	nc, err := nats.Connect(cfg.Nats.Addr)
	if err != nil {
		log.Fatal(err)
	}

	return nc
}

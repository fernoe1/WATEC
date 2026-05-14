package nats

import (
	"log"

	"github.com/nats-io/nats.go"
)

func NewNATS() *nats.Conn {
	nc, err := nats.Connect("nats://nats:4222")
	if err != nil {
		log.Fatal(err)
	}

	return nc
}

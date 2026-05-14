package notification

import (
	"context"
	"errors"
	"time"

	"github.com/fernoe1/WATEC/notification/config"
	"github.com/nats-io/nats.go"
)

func Subscribe(ctx context.Context, nc *nats.Conn, subject string, handler Handler, cfg *config.Config) error {
	subj := cfg.NotificationSubject + "." + subject

	sub, err := nc.SubscribeSync(subj)

	if err != nil {
		return err
	}
	defer sub.Unsubscribe()

	for {
		msg, err := sub.NextMsg(7 * time.Second)
		if errors.Is(err, nats.ErrTimeout) {
			continue
		}

		if err != nil {
			return err
		}

		if err := handler(ctx, msg); err != nil {
			return err
		}
	}
}

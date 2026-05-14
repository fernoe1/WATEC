package notification

import (
	"context"
	"errors"
	"time"

	"github.com/fernoe1/WATEC/notification/config"
	"github.com/nats-io/nats.go"
)

type Subscriber struct {
	nc  *nats.Conn
	cfg *config.Config
}

func NewNotificationSubscriber(
	nc *nats.Conn,
	cfg *config.Config,
) *Subscriber {
	return &Subscriber{nc: nc, cfg: cfg}
}

func (ns *Subscriber) Subscribe(
	ctx context.Context,
	subj string,
	handler Handler,
) error {
	subj = ns.cfg.Nats.NotificationSubject + "." + subj

	sub, err := ns.nc.SubscribeSync(subj)

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

package notification

import (
	"context"

	"github.com/fernoe1/WATEC/gateway/config"
	notifpb "github.com/fernoe1/protogen/watec/shares/notification"
	"github.com/nats-io/nats.go"
	"google.golang.org/protobuf/proto"
)

type Publisher struct {
	nc  *nats.Conn
	cfg *config.Config
}

func NewNotificationPublisher(
	nc *nats.Conn,
	cfg *config.Config,
) *Publisher {
	return &Publisher{nc: nc, cfg: cfg}
}

func (np *Publisher) Publish(
	ctx context.Context,
	subj string,
	notif *notifpb.Notification,
) error {
	raw, err := proto.Marshal(notif)
	if err != nil {
		return err
	}

	err = np.nc.Publish(np.cfg.NotificationSubject+"."+subj, raw)
	if err != nil {
		return err
	}

	return nil
}

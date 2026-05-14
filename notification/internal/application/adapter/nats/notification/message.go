package notification

import (
	"github.com/fernoe1/WATEC/notification/internal/application/domain"
	notifpb "github.com/fernoe1/protogen/watec/shares/notification"
	"github.com/nats-io/nats.go"
	"google.golang.org/protobuf/proto"
)

func ToDomain(msg *nats.Msg) (*domain.Notification, error) {
	var pb notifpb.Notification
	if err := proto.Unmarshal(msg.Data, &pb); err != nil {
		return nil, err
	}

	return &domain.Notification{Email: pb.Email}, nil
}

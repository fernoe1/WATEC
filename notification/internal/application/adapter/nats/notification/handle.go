package notification

import (
	"context"

	"github.com/fernoe1/WATEC/notification/internal/application"
	"github.com/nats-io/nats.go"
)

type Handler func(context.Context, *nats.Msg) error

type MailjetHandler struct {
	uc application.NotificationUsecase
}

func NewMailjetHandler(uc application.NotificationUsecase) *MailjetHandler {
	return &MailjetHandler{uc: uc}
}

func (h *MailjetHandler) Handler(ctx context.Context, msg *nats.Msg) error {
	notif, err := ToDomain(msg)
	if err != nil {

		return err
	}

	err = h.uc.Send(ctx, notif)
	if err != nil {

		return err
	}

	return nil
}

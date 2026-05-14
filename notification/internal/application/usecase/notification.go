package usecase

import (
	"context"

	"github.com/fernoe1/WATEC/notification/internal/application"
	"github.com/fernoe1/WATEC/notification/internal/application/domain"
)

type NotificationUsecase struct {
	p application.Presenter
}

func NewNotificationUsecase(p application.Presenter) *NotificationUsecase {
	return &NotificationUsecase{p: p}
}

func (n *NotificationUsecase) Send(ctx context.Context, notification *domain.Notification) error {
	return n.p.Send(ctx, notification)
}

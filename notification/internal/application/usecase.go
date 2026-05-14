package application

import (
	"context"

	"github.com/fernoe1/WATEC/notification/internal/application/domain"
)

type Presenter interface {
	Send(ctx context.Context, notification *domain.Notification) error
}

type NotificationUsecase interface {
	Send(ctx context.Context, notification *domain.Notification) error
}

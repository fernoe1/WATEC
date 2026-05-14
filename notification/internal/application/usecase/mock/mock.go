package mock

import (
	"context"

	"github.com/fernoe1/WATEC/notification/internal/application/domain"
	"github.com/stretchr/testify/mock"
)

type Mailjet struct{ mock.Mock }

func (m *Mailjet) Send(ctx context.Context, notification *domain.Notification) error {
	args := m.Called(ctx, notification)
	return args.Error(0)
}

package usecase

import (
	"context"
	"testing"

	"github.com/fernoe1/WATEC/notification/internal/application/domain"
	MOCK "github.com/fernoe1/WATEC/notification/internal/application/usecase/mock"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestNotificationUsecase_Send(T *testing.T) {
	table := []struct {
		name    string
		prepare func(*MOCK.Mailjet)
		wantErr bool
	}{
		{
			name: "success",
			prepare: func(m *MOCK.Mailjet) {
				m.On("Send", mock.Anything, mock.AnythingOfType("*domain.Notification")).Return(nil)
			},
		},
	}

	for _, t := range table {
		T.Run(t.name, func(T *testing.T) {
			m := new(MOCK.Mailjet)
			t.prepare(m)

			nc := NotificationUsecase{
				p: m,
			}

			err := nc.Send(context.Background(), &domain.Notification{Email: "krutoytemirlan2007@gmail.com"})

			if t.wantErr {
				require.Error(T, err)
			} else {
				require.NoError(T, err)
			}

			m.AssertExpectations(T)
		})
	}
}

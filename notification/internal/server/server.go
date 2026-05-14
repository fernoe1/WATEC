package server

import (
	"context"

	"github.com/fernoe1/WATEC/notification/config"
	"github.com/fernoe1/WATEC/notification/internal/application/adapter/mailjet"
	"github.com/fernoe1/WATEC/notification/internal/application/adapter/nats/notification"
	"github.com/fernoe1/WATEC/notification/internal/application/usecase"
	"github.com/nats-io/nats.go"
)

type Server struct {
	nc  *nats.Conn
	cfg *config.Config
}

func NewServer(nc *nats.Conn, cfg *config.Config) *Server {
	return &Server{nc: nc, cfg: cfg}
}

func (s *Server) Run() error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	presenter := mailjet.NewMailjetPresenter(s.cfg)
	notificationUC := usecase.NewNotificationUsecase(presenter)

	mailjet := notification.NewMailjetHandler(notificationUC)

	notificationSubscriber := notification.NewNotificationSubscriber(s.nc, s.cfg)

	err := notificationSubscriber.Subscribe(ctx, "mailjet", mailjet.Handler)
	if err != nil {
		return err
	}

	return nil
}

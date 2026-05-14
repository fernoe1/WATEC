package mailjet

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/fernoe1/WATEC/notification/config"
	"github.com/fernoe1/WATEC/notification/internal/application/domain"
)

type MailjetPresenter struct {
	apiKey    string
	apiSecret string
	fromEmail string
	fromName  string
	client    *http.Client
}

func NewMailjetPresenter(cfg *config.Config) *MailjetPresenter {
	return &MailjetPresenter{
		apiKey:    cfg.Mailjet.ApiKey,
		apiSecret: cfg.Mailjet.ApiSecret,
		fromEmail: cfg.Mailjet.FromEmail,
		fromName:  cfg.Mailjet.FromName,
		client:    &http.Client{Timeout: 5 * time.Second},
	}
}

type sendPayload struct {
	Messages []struct {
		From struct {
			Email string `json:"Email"`
			Name  string `json:"Name"`
		} `json:"From"`
		To []struct {
			Email string `json:"Email"`
			Name  string `json:"Name,omitempty"`
		} `json:"To"`
		Subject  string `json:"Subject"`
		TextPart string `json:"TextPart"`
	} `json:"Messages"`
}

type sendResponse struct {
	Messages []struct {
		To []struct {
			MessageUUID string `json:"MessageUUID"`
			MessageID   int64  `json:"MessageID"`
		} `json:"To"`
	} `json:"Messages"`
}

func (m *MailjetPresenter) generatePayload(notification *domain.Notification) sendPayload {
	return sendPayload{
		Messages: []struct {
			From struct {
				Email string `json:"Email"`
				Name  string `json:"Name"`
			} `json:"From"`
			To []struct {
				Email string `json:"Email"`
				Name  string `json:"Name,omitempty"`
			} `json:"To"`
			Subject  string `json:"Subject"`
			TextPart string `json:"TextPart"`
		}{
			{
				From: struct {
					Email string `json:"Email"`
					Name  string `json:"Name"`
				}{
					Email: m.fromEmail,
					Name:  m.fromName,
				},
				To: []struct {
					Email string `json:"Email"`
					Name  string `json:"Name,omitempty"`
				}{{Email: notification.Email}},
				Subject:  "subject",
				TextPart: "text part",
			},
		},
	}
}

func (m *MailjetPresenter) Send(ctx context.Context, notification *domain.Notification) error {
	in := m.generatePayload(notification)

	raw, err := json.Marshal(in)
	if err != nil {
		return err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, "https://api.mailjet.com/v3.1/send", bytes.NewReader(raw))
	if err != nil {
		return err
	}

	req.SetBasicAuth(m.apiKey, m.apiSecret)
	req.Header.Set("Content-Type", "application/json")

	res, err := m.client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode >= 500 {
		return fmt.Errorf("%d", res.StatusCode)
	}
	if res.StatusCode >= 400 {
		return fmt.Errorf("%d", res.StatusCode)
	}

	var out sendResponse
	if err := json.NewDecoder(res.Body).Decode(&out); err != nil {
		return err
	}

	slog.Info("response", "out", out)

	return nil
}

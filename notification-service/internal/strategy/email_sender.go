package strategy

import (
	"errors"
	"notification-service/internal/domain"
)

type EmailProvider interface {
	SendEmail(domain.Notification) error
}

type EmailSender struct {
	providers []EmailProvider
}

func NewEmailSender(p []EmailProvider) *EmailSender {
	return &EmailSender{providers: p}
}

func (e *EmailSender) Send(n domain.Notification) error {
	for _, p := range e.providers {
		if err := p.SendEmail(n); err == nil {
			return nil
		}
	}
	return errors.New("all email providers failed")
}
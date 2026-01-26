package email

import (
	"errors"
	"notification-service/internal/domain"
)

type SMTPProvider struct{}

func (s *SMTPProvider) SendEmail(email domain.Notification) error {
	return errors.New("smtp down")
}
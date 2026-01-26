package email

import (
	"fmt"
	"notification-service/internal/domain"
)

type SendGridProvider struct{}

func (s *SendGridProvider) SendEmail(n domain.Notification) error {
	fmt.Println("SendGrid sent email to:", n.To)
	return nil
}

package strategy

import "notification-service/internal/domain"

type Sender interface {
	Send(domain.Notification) error
}
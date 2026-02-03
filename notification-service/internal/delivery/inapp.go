package delivery

import (
	"log"
	"notification-service/internal/domain"
)

type InAppSender struct{}

func (i *InAppSender) SendInAppNotification(event domain.NotificationEvent) {
	log.Printf("we have a email in app push %s", event.Title)
}
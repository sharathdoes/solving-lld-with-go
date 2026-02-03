package notifications

import (
	"log"
	"notification-service/internal/delivery"
	 "notification-service/internal/domain"
)

type NotificationProcessor struct {
	InAppSender delivery.InAppSender
	EmailSender delivery.EmailSender
}

func (p *NotificationProcessor) Process(event domain.NotificationEvent) {

	log.Printf("Started working on Notification %s", event.Title)
	switch event.TypeofN {
	case domain.InApp:
		p.InAppSender.SendInAppNotification(event)
	case domain.Email:
		p.EmailSender.Send(event)
	default:
		log.Printf("Can't send notifications of that type")
	}
}

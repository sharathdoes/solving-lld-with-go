package notifications

import (
	"log"
	"notification-service/internal/events"
)
func StartNotificationsWorker(processor NotificationProcessor){
	go func() {
		for event := range events.NotificationChannel {
			log.Print("Received from channel ")
			processor.Process(event)
		}
	}()
}
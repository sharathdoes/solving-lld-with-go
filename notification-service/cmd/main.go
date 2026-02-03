package main

import (
	"notification-service/internal/api"
	"notification-service/internal/delivery"
	"notification-service/internal/notifications"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	
	processor:=notifications.NotificationProcessor{
		InAppSender: delivery.InAppSender{},
		EmailSender: delivery.EmailSender{
			SMTPHost: "smtp.gmail.com",
			SMTPPort: "587",
			Username: "justabountyhunter935@gmail.com",
			Password: "cvnnegxwyvxoxdom",
			From:     "justabountyhunter935@gmail.com",
		},
	}

	notifications.StartNotificationsWorker(processor)
	api.RegisterRoutes(r)
	r.Run(":8080")
}

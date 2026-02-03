package main

import (
	"log"
	"notification-service/internal/api"
	"notification-service/internal/delivery"
	"notification-service/internal/notifications"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
    if err != nil {
        log.Println("No .env file found, fetching from system env")
    }
	r := gin.Default()
	
	processor:=notifications.NotificationProcessor{
		InAppSender: delivery.InAppSender{},
		EmailSender: delivery.EmailSender{
			SMTPHost: "smtp.gmail.com",
			SMTPPort: "587",
			Username: "justabountyhunter935@gmail.com",
			Password: os.Getenv("PASSWORD"),
			From:     "justabountyhunter935@gmail.com",
		},
	}

	notifications.StartNotificationsWorker(processor)
	api.RegisterRoutes(r)
	r.Run(":8080")
}

package events

import (
	"notification-service/internal/domain"
)

var NotificationChannel = make(chan domain.NotificationEvent, 100)


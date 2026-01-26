package queue

import (
	"notification-service/internal/domain"
)

type Queue struct {
	Jobs chan domain.Notification
}

func NewQueue(size int) *Queue {
	jobs:=make(chan domain.Notification, size)
	return &Queue{jobs}
}
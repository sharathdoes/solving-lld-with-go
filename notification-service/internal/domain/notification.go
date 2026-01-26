package domain

type NotificationType string

const (
	Email NotificationType="email"
	SMS   NotificationType = "sms"
	Push  NotificationType = "push"
)


type Notification struct {
	ID      string
	Type    NotificationType
	To      string
	Title   string
	Message string
	Retries int
}
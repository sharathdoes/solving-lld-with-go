package domain

type NotificationType string

const (
	InApp NotificationType = "IN_APP"
	Email NotificationType = "EMAIL"
)

type NotificationEvent struct {
	UserId string
	TypeofN NotificationType
	Title string
	Message string
	Email string
}
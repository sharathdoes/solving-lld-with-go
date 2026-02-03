package api

import (
	"net/http"
	"notification-service/internal/domain"
	"notification-service/internal/events"

	"github.com/gin-gonic/gin"
)

type CreateNotificationRequest struct {
	UserID string `json:"user_id"`
	Type string `json:"type"`
	Title string `json:"title"`
	Message string `json:"message"`
	Email string `json:"email"`
}


func createNotificion(c *gin.Context){
	var req CreateNotificationRequest
	if err:=c.ShouldBindJSON(&req); err!=nil {
		c.JSON(http.StatusBadRequest, gin.H{"Body doesn't match": err.Error()})
		return
	}

	if req.Type == string(domain.Email) && req.Email=="" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "email is required for EMAIL notification",
		})
		return
	}
	
	events.NotificationChannel <- domain.NotificationEvent{
		UserId:  req.UserID,
		TypeofN:    domain.NotificationType(req.Type),
		Title:   req.Title,
		Message: req.Message,
		Email:   req.Email,
	}

	c.JSON(http.StatusAccepted, gin.H{"status": "queued"})
}
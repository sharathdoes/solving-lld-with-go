package http

import (
	"rate-limiter/internal/middleware"
	"rate-limiter/internal/ratelimit"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, l ratelimit.Limiter) {
	r.Use(middleware.Ratelimit(l))
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message":"pong"})
	})
}
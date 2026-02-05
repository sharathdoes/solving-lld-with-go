package middleware

import (
	"net/http"

	"rate-limiter/internal/ratelimit"

	"github.com/gin-gonic/gin"
)

func resolveKey(c *gin.Context) string {
	if userID, exists := c.Get("userID"); exists {
		return "user:" + userID.(string)
	}
	return "ip:" + c.ClientIP()
}


func Ratelimit( l ratelimit.Limiter) gin.HandlerFunc {
	return func(c *gin.Context){
		key:=resolveKey(c)
		allowed, err:=l.Allow(c.Request.Context(), key)
		if err !=nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error":"rate limit error"})
			return
		}

		if !allowed {
			c.Header("Retry-After", "1")
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"error": "rate limit exceeded",
			})
			return
		}

		c.Next()
	}
}


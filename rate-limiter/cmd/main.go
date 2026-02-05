package main

import (
	"github.com/gin-gonic/gin"
	"rate-limiter/internal/http"
	"rate-limiter/internal/ratelimit"
)

func main() {
	r := gin.Default()

	l := ratelimit.NewMemoryLimiter(
		1,   // capacity
		1.0,  // tokens per second
	)

	http.RegisterRoutes(r, l)

	r.Run(":8080")
}

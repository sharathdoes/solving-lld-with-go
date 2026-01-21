package auth

import (
	"simple-todo/internal/config"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoutes(r *gin.Engine, db *gorm.DB, cfg *config.Config) {
	repo := NewRepository(db)
	svc := NewService(repo, cfg.JWTSecret, cfg.AccessTokenTTL, cfg.RefreshTokenTTL)
	h := NewHandler(svc)

	g := r.Group("/auth")
	{
		g.POST("/signup", h.SignUp)
		g.POST("/login", h.Login)
		g.POST("/refresh", h.Refresh)
	}
}

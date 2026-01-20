package server

import (
	"user-auth-service/internal/config"
	"user-auth-service/internal/modules/auth"
	"user-auth-service/internal/modules/health"

	"user-auth-service/internal/database"

	"github.com/gin-gonic/gin"
)

type Server struct {
	engine *gin.Engine
	config *config.Config
}

func New(cfg *config.Config) *Server {
	r := gin.Default()
	database, err := database.Connect(cfg.DBUrl)
	if err != nil {
		panic(err)
	}
	auth.RegisterRoutes(r, database, cfg)
	r.GET("/ping", modules.Health)
	
	return &Server{
		engine: r,
		config: cfg,
	}
}

func (s *Server) Run() error {
	return s.engine.Run(":" + s.config.Port)
}

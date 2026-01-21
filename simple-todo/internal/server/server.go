package server

import (
	"log"
	"simple-todo/internal/config"
	"simple-todo/internal/database"

	"github.com/gin-gonic/gin"
)

type Server struct {
	engine *gin.Engine
	config *config.Config
}

func NewServer(c *config.Config) *Server {
	r := gin.Default()
	_, err := database.Connect(c.DBUrl)
	if err != nil {
		panic("Database Connection Failed")
	} else {
		log.Default().Println("Database Connected !!")
	}
	return &Server{engine: r, config: c}
}

func (s *Server) Run() error {
	return s.engine.Run(":" + s.config.Port)
}

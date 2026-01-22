package server

import (
	"log"
	"simple-todo/internal/config"
	"simple-todo/internal/database"
	"simple-todo/internal/modules/auth"
	"simple-todo/internal/modules/projects"
	"simple-todo/internal/modules/tasks"

	"github.com/gin-gonic/gin"
)

type Server struct {
	engine *gin.Engine
	config *config.Config
}

func NewServer(c *config.Config) *Server {
	r := gin.Default()
	database, err := database.Connect(c.DBUrl)
	auth.RegisterRoutes(r, database, c)
	tasks.TaskRoutes(r,database,c)
	projects.ProjectRoutes(r,database,c)
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

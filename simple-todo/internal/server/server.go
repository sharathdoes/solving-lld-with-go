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
	db, err := database.Connect(c.DBUrl)
	database.RunMigrations(c.DBUrl)
	auth.RegisterRoutes(r, db, c)
	tasks.TaskRoutes(r,db,c)
	projects.ProjectRoutes(r,db,c)
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

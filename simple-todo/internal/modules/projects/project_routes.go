package projects

import (
	"simple-todo/internal/config"
	"simple-todo/internal/modules/auth"
	"simple-todo/internal/modules/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)


func ProjectRoutes(r *gin.Engine, db *gorm.DB, cfg *config.Config){
	repo:=NewRepository(db)
	userRepo:=auth.NewRepository(db);
	svc:=NewService(repo, *userRepo)
	h:=NewHandler(svc)
	
	g:=r.Group("/project")
	g.GET("/getAll", h.GetProjects)
	// g.GET("/:id", h.FindByID)
	g.GET("/getById",h.FindByID)
	g.Use(middleware.AuthMiddleware(cfg.JWTSecret))

	g.POST("/create", h.CreateProject)
	g.POST("/update", h.UpdateProject)
	
}
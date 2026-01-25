package tasks


import (
	"simple-todo/internal/config"
	"simple-todo/internal/modules/auth"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)


func TaskRoutes(r *gin.Engine, db *gorm.DB, cfg *config.Config){
	repo:=NewRepository(db)
	userRepo:=auth.NewRepository(db);
	svc:=NewService(repo, userRepo)
	h:=NewHandler(svc)
	
	g:=r.Group("/task")
	{	
	g.POST("/create", h.CreateTask)
	g.GET("/getAll", h.GetTasks)
	}
	
}
package projects

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	srv *Service
}

func NewHandler(srv *Service) *Handler {
	return &Handler{srv: srv}
}

func (h *Handler) CreateProject(c *gin.Context ){
	var project CreateProjectDTO
	if err:=c.ShouldBindBodyWithJSON(&project); err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	proj, err:=h.srv.CreateProject(c, project.Title, project.Description, project.OwnerID, project.MemberIDs,)
	if err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, proj)
}

func(h )
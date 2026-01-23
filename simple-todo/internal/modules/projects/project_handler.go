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

func (h *Handler) CreateProject(c *gin.Context) {
	var project CreateProjectDTO
	if err := c.ShouldBindJSON(&project); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	userIDRaw, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "UNAUTHORIZED"})
		return
	}
	OwnerID := userIDRaw.(string)
	proj, err := h.srv.CreateProject(c, project.Title, project.Description, OwnerID, project.MemberIDs)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, proj)
}

func (h *Handler) UpdateProject(c *gin.Context) {
	var body UpdateProjectInput
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	userIDRaw, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "UNAUTHORIZED"})
		return
	}

	OwnerID := userIDRaw.(string)
	proj, err := h.srv.UpdateProject(c, body.ID, body.Title, body.Description, OwnerID, body.MemberIDs)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, proj)
}

func (h *Handler) GetProjects(c *gin.Context) {
	proj, err := h.srv.GetProjects(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, proj)
}

func (h *Handler) FindByID(c *gin.Context) {
	id := c.Query("id")
	proj, err := h.srv.repo.FindById(c,id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, proj)
}

func (h *Handler) FindMyProjects(c *gin.Context) {
	id := c.Query("id")
	proj, err := h.srv.repo.FindMyProjects(c,id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, proj)
}

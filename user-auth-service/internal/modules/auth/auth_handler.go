package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
  svc *Service
}

func NewHandler(s *Service) *Handler {
	return &Handler{s}
}

func (h * Handler) SignUp(c *gin.Context){
	var body SignUpRequest
	if err:=c.ShouldBindJSON(&body); err!=nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if err:=h.svc.SignUp(c,body.Username,body.Email, body.Password); err!=nil {
		c.JSON(http.StatusForbidden, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message":"user created"})	
}


func (h *Handler) Login(c *gin.Context) {
  var req LoginRequest
  if err := c.ShouldBindJSON(&req); err != nil {
    c.JSON(400, gin.H{"error": err.Error()})
    return
  }

  access, refresh, err := h.svc.Login(c, req.Email, req.Password)
  if err != nil {
    c.JSON(401, gin.H{"error": "invalid credentials"})
    return
  }

  c.JSON(200, gin.H{
    "access_token":  access,
    "refresh_token": refresh,
  })
}

func (h *Handler) Refresh(c *gin.Context) {
  var req RefreshRequest
  if err := c.ShouldBindJSON(&req); err != nil {
    c.JSON(400, gin.H{"error": err.Error()})
    return
  }

  access, refresh, err := h.svc.Refresh(c, req.RefreshToken)
  if err != nil {
    c.JSON(401, gin.H{"error": "invalid refresh token"})
    return
  }

  c.JSON(200, gin.H{
    "access_token":  access,
    "refresh_token": refresh,
  })
}

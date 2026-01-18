package handlers

import (
	"gin-quickstart/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UrlHandler struct {
	service *service.UrlService
}

func NewUrlHander(service *service.UrlService) *UrlHandler {
		return &UrlHandler{service: service}
}

//These two functions are your HTTP endpoints (controllers / route handlers).

func (h *UrlHandler) ShorteURL( c *gin.Context){
	var body struct {
		URL string `json:"url"`
	}
	if err:=c.BindJSON(&body); err!=nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":"invalid request"})
		return
	}
	code:=h.service.ShortenUrl(body.URL)
	c.JSON(http.StatusOK, gin.H{"short_url": code})
	
}

func (h *UrlHandler) Redirect ( c *gin.Context){
	code:=c.Param("code")
	url, ok:=h.service.Resolve(code);
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}
	c.Redirect(http.StatusFound, url)
}
package main

import (
	"gin-quickstart/store"
	"gin-quickstart/service"
	"gin-quickstart/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
 	r := gin.Default()

  	store:=store.NewMemoryStore()
  	service := service.NewUrlService(store) //logic
	handler := handlers.NewUrlHander(service) //HTTP, JSON, status codes

	r.POST("/shorten", handler.ShorteURL) 
	r.GET("/:code", handler.Redirect)
    r.GET("/ping", func(c *gin.Context) {
    c.JSON(200, gin.H{
      "message": "pong",
    })
  })
  r.Run() 
}
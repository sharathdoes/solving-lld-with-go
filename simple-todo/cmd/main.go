package main

import (
	"log"
	"simple-todo/internal/config"
	"simple-todo/internal/server"
)

func main() {

	c:=config.Load()
	srv:=server.NewServer(c)
	log.Fatal(srv.Run())

}
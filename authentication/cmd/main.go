package main

import (
	"log"
	"user-auth-service/internal/config"
	"user-auth-service/internal/server"
)

func main() {
	cfg := config.Load()
	srv := server.New(cfg)
	log.Fatal(srv.Run())

}

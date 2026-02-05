package main

import (
	"log"
	"authentication/internal/config"
	"authentication/internal/server"
)

func main() {
	cfg := config.Load()
	srv := server.New(cfg)
	log.Fatal(srv.Run())

}

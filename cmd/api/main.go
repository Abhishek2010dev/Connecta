package main

import (
	"log"

	"github.com/Abhishek2010dev/Connecta/internal/server"
	"github.com/Abhishek2010dev/Connecta/pkg/config"
)

func main() {
	cfg := config.Load()
	server := server.New(cfg)
	log.Printf("Server started at %s:%s", cfg.Server.Host, cfg.Server.Port)
	if err := server.ListenAndServe(); err != nil {
		log.Fatal("Failed to started server:", err)
	}
}

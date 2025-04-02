package main

import (
	"log"

	"github.com/Abhishek2010dev/Connecta/internal/cache"
	"github.com/Abhishek2010dev/Connecta/internal/database"
	"github.com/Abhishek2010dev/Connecta/internal/server"
	"github.com/Abhishek2010dev/Connecta/pkg/config"
)

func main() {
	cfg := config.Load()

	database, err := database.New(cfg.Database)
	if err != nil {
		log.Fatal(err)
	}
	defer database.Close()

	cache, err := cache.New(cfg.Redis)
	if err != nil {
		log.Fatal(err)
	}
	defer cache.Close()

	server := server.New(cfg, database.Get(), cache.Get())
	log.Printf("Server is running at %s:%s", cfg.Server.Host, cfg.Server.Port)
	if err := server.ListenAndServe(); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}

package main

import (
	"log"

	"github.com/Abhishek2010dev/Connecta/internal/server"
	"github.com/Abhishek2010dev/Connecta/pkg/config"
)

func main() {
	server := server.New(config.Load())
	if err := server.ListenAndServe(); err != nil {
		log.Fatal("Failed to started server:", err)
	}
}

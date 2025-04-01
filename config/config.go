package config

import (
	"log"
	"os"
)

type ServerConfig struct {
	Host string
	Port string
}

type Config struct {
	server ServerConfig
}

func getEnv(key string) {
	value, exits := os.LookupEnv(key)
	if value == "" || !exits {
		log.Fatalf("Error: Missing environment variable %s", key)
	}
}

func Load() Config {
}

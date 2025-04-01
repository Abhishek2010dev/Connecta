package config

import (
	"log"
	"time"
)

type ServerConfig struct {
	Host    string
	Port    string
	Timeout time.Duration
}

func NewServerConfig() ServerConfig {
	timeout, err := time.ParseDuration(getEnv("SERVER_TIMEOUT"))
	if err != nil {
		log.Fatal("Failed to parse timeout")
	}

	return ServerConfig{
		Host:    getEnv("SERVER_HOST"),
		Port:    getEnv("SERVER_PORT"),
		Timeout: timeout,
	}
}

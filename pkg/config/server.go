package config

import (
	"log"
	"time"
)

type Server struct {
	Host    string
	Port    string
	Timeout time.Duration
}

func NewServerConfig() Server {
	timeout, err := time.ParseDuration(getEnv("SERVER_TIMEOUT"))
	if err != nil {
		log.Fatal("Failed to parse timeout")
	}

	return Server{
		Host:    getEnv("SERVER_HOST"),
		Port:    getEnv("SERVER_PORT"),
		Timeout: timeout,
	}
}

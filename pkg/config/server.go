package config

import (
	"time"
)

type Server struct {
	Host    string
	Port    string
	Timeout time.Duration
}

func NewServer() Server {
	return Server{
		Host:    LoadEnv("SERVER_HOST"),
		Port:    LoadEnv("SERVER_PORT"),
		Timeout: LoadEnvDuration("SERVER_TIMEOUT"),
	}
}

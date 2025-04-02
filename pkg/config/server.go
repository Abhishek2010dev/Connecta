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
		Host:    getEnv("SERVER_HOST"),
		Port:    getEnv("SERVER_PORT"),
		Timeout: getDurationEnv("SERVER_TIMEOUT"),
	}
}

package config

import (
	"time"
)

type Server struct {
	Host       string
	Port       string
	CsrfSecure string
	Timeout    time.Duration
}

func NewServer() Server {
	return Server{
		Host:       LoadEnv("SERVER_HOST"),
		Port:       LoadEnv("SERVER_PORT"),
		CsrfSecure: LoadEnv("CSRF_SECRET"),
		Timeout:    LoadEnvDuration("SERVER_TIMEOUT"),
	}
}

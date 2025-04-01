package server

import (
	"fmt"
	"net/http"

	"github.com/Abhishek2010dev/Connecta/pkg/config"
)

type Server struct{}

func New(cfg config.Config) *http.Server {
	NewServer := Server{}
	server := &http.Server{
		Addr:         fmt.Sprintf("%s:%s", cfg.Server.Host, cfg.Server.Port),
		Handler:      NewServer.RegisterRoutes(),
		ReadTimeout:  cfg.Server.Timeout,
		WriteTimeout: cfg.Server.Timeout,
	}
	return server
}

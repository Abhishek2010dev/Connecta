package server

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/Abhishek2010dev/Connecta/pkg/config"
)

type Server struct {
	db *sql.DB
}

func New(cfg config.Config, db *sql.DB) *http.Server {
	NewServer := Server{
		db: db,
	}

	server := &http.Server{
		Addr:         fmt.Sprintf("%s:%s", cfg.Server.Host, cfg.Server.Port),
		Handler:      NewServer.RegisterRoutes(),
		ReadTimeout:  cfg.Server.Timeout,
		WriteTimeout: cfg.Server.Timeout,
	}
	return server
}

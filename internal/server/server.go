package server

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/Abhishek2010dev/Go-Htmx-Auth-Example/pkg/config"
)

type Server struct {
	db  *sql.DB
	cfg config.Config
}

func New(cfg config.Config, db *sql.DB) *http.Server {
	NewServer := Server{
		db:  db,
		cfg: cfg,
	}

	server := &http.Server{
		Addr:         fmt.Sprintf("%s:%s", cfg.Server.Host, cfg.Server.Port),
		Handler:      NewServer.RegisterRoutes(),
		ReadTimeout:  cfg.Server.Timeout,
		WriteTimeout: cfg.Server.Timeout,
	}
	return server
}

package server

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/Abhishek2010dev/Connecta/pkg/config"
	"github.com/redis/go-redis/v9"
)

type Server struct {
	db    *sql.DB
	cache *redis.Client
	cfg   config.Config
}

func New(cfg config.Config, db *sql.DB, cache *redis.Client) *http.Server {
	NewServer := Server{
		db:    db,
		cache: cache,
	}

	server := &http.Server{
		Addr:         fmt.Sprintf("%s:%s", cfg.Server.Host, cfg.Server.Port),
		Handler:      NewServer.RegisterRoutes(),
		ReadTimeout:  cfg.Server.Timeout,
		WriteTimeout: cfg.Server.Timeout,
	}
	return server
}

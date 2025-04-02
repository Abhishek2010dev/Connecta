package server

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/Abhishek2010dev/Connecta/internal/database"
	"github.com/Abhishek2010dev/Connecta/pkg/config"
)

type Server struct {
	db *sql.DB
}

func New(cfg config.Config) *http.Server {
	database, err := database.New(cfg.Database)
	if err != nil {
		log.Fatal(err)
	}
	defer database.Close()

	NewServer := Server{
		db: database.Get(),
	}

	server := &http.Server{
		Addr:         fmt.Sprintf("%s:%s", cfg.Server.Host, cfg.Server.Port),
		Handler:      NewServer.RegisterRoutes(),
		ReadTimeout:  cfg.Server.Timeout,
		WriteTimeout: cfg.Server.Timeout,
	}
	return server
}

package database

import (
	"database/sql"
	"fmt"

	"github.com/Abhishek2010dev/Connecta/pkg/config"
)

type postgresqlDatabase struct {
	db *sql.DB
}

func New(cfg config.Database) (Provider, error) {
	db, err := sql.Open("postgres", cfg.URL)
	if err != nil {
		return nil, fmt.Errorf("Failed to create connection: %w", err)
	}

	db.SetConnMaxIdleTime(cfg.ConnMaxIdleTime)
	db.SetConnMaxLifetime(cfg.ConnMaxLifetime)
	db.SetMaxIdleConns(cfg.MaxIdleConns)
	db.SetMaxOpenConns(cfg.MaxOpenConns)

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("Failed to ping connection: %w", err)
	}

	return &postgresqlDatabase{db: db}, nil
}

func (p *postgresqlDatabase) Get() *sql.DB {
	return p.db
}

func (p *postgresqlDatabase) Close() error {
	panic("not implemented") // TODO: Implement
}

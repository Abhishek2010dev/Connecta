package database

import "database/sql"

type Provider interface {
	Get() *sql.DB
	Close() error
}

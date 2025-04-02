package database

import "database/sql"

type DBProvider interface {
	Get() *sql.DB
	Close() error
}

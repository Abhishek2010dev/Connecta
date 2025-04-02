package config

import "time"

type Database struct {
	URL             string
	MaxOpenConns    int
	MaxIdleConns    int
	ConnMaxLifetime time.Duration
	ConnMaxIdleTime time.Duration
}

func NewDatabase() Database {
	return Database{
		URL: getEnv("DATABASE_URL"),
	}
}

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
		URL:             LoadEnv("DB_URL"),
		MaxOpenConns:    LoadEnvInt("DB_MAX_OPEN_CONNS"),
		MaxIdleConns:    LoadEnvInt("DB_MAX_IDLE_CONNS"),
		ConnMaxLifetime: LoadEnvDuration("DB_CONN_MAX_LIFETIME"),
		ConnMaxIdleTime: LoadEnvDuration("DB_CONN_MAX_IDLE_TIME"),
	}
}

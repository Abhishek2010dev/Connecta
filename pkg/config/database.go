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
		URL:             getEnv("DB_URL"),
		MaxOpenConns:    getIntEnv("DB_MAX_OPEN_CONN"),
		MaxIdleConns:    getIntEnv("DB_MAX_IDLE_CONN"),
		ConnMaxLifetime: getDurationEnv("DB_CONN_MAX_LIFETIME"),
		ConnMaxIdleTime: getDurationEnv("DB_CONN_MAX_IDLE_TIME"),
	}
}

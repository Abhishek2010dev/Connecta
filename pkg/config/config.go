package config

import _ "github.com/joho/godotenv/autoload"

type Config struct {
	Server
	Database
	Redis
}

func Load() Config {
	return Config{
		Server:   NewServer(),
		Database: NewDatabase(),
		Redis:    NewRedis(),
	}
}

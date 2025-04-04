package config

import _ "github.com/joho/godotenv/autoload"

type Config struct {
	Server
	Database
	Redis
	Auth
}

func Load() Config {
	return Config{
		Server:   NewServer(),
		Database: NewDatabase(),
		Redis:    NewRedis(),
		Auth:     NewAuth(),
	}
}

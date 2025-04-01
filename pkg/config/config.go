package config

import _ "github.com/joho/godotenv/autoload"

type Config struct {
	Server
	Database
}

func Load() Config {
	return Config{
		Server:   NewServerConfig(),
		Database: NewDatabase(),
	}
}

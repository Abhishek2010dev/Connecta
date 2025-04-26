package config

import _ "github.com/joho/godotenv/autoload"

type Config struct {
	Server
	Database
	Auth
}

func Load() Config {
	return Config{
		Server:   NewServer(),
		Database: NewDatabase(),
		Auth:     NewAuth(),
	}
}

package config

import _ "github.com/joho/godotenv/autoload"

type Config struct {
	server ServerConfig
}

func Load() Config {
	return Config{
		server: NewServerConfig(),
	}
}

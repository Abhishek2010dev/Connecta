package config

import _ "github.com/joho/godotenv/autoload"

type Config struct {
	Server ServerConfig
}

func Load() Config {
	return Config{
		Server: NewServerConfig(),
	}
}

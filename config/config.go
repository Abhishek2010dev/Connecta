package config

type ServerConfig struct {
	Host string
	Port string
}

type Config struct {
	server ServerConfig
}

func Load() Config {
	return Config{
		server: ServerConfig{
			Host: getEnv("SERVER_HOST"),
			Port: getEnv("SERVER_PORT"),
		},
	}
}

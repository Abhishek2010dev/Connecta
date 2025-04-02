package config

type Redis struct {
	URL string
}

func NewRedis() Redis {
	return Redis{
		URL: getEnv("REDIS_URL"),
	}
}

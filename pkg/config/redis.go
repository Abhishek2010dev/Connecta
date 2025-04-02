package config

type Redis struct {
	URL string
}

func NewRedis() Redis {
	return Redis{
		URL: LoadEnv("REDIS_URL"),
	}
}

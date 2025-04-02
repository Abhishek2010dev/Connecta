package cache

import (
	"github.com/Abhishek2010dev/Connecta/pkg/config"
	"github.com/redis/go-redis/v9"
)

type redisClient struct {
	client *redis.Client
}

func New(cfg config.Redis) Provider {
	client := redis.NewClient(&redis.Options{
		Addr: cfg.URL,
	})
	return &redisClient{client}
}

func (c *redisClient) Get() *redis.Client {
	return c.client
}

func (c *redisClient) Close() error {
	return c.Close()
}

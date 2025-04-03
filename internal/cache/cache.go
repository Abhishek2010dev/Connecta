package cache

import (
	"fmt"
	"log"

	"github.com/Abhishek2010dev/Connecta/pkg/config"
	"github.com/redis/go-redis/v9"
)

type redisClient struct {
	client *redis.Client
}

func New(cfg config.Redis) (Provider, error) {
	opts, err := redis.ParseURL(cfg.URL)
	if err != nil {
		return nil, fmt.Errorf("Failed to parse url: %w", err)
	}
	log.Println("Successfully connected to redis")
	return &redisClient{client: redis.NewClient(opts)}, nil
}

func (c *redisClient) Get() *redis.Client {
	return c.client
}

func (c *redisClient) Close() error {
	return c.Close()
}

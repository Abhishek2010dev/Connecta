package cache

import "github.com/redis/go-redis/v9"

type redisClient struct {
	client *redis.Client
}

func (c *redisClient) Get() *redis.Client {
	return c.client
}

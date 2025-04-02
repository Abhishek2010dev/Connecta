package cache

import (
	"github.com/redis/go-redis/v9"
)

type Provider interface {
	Get() *redis.Client
	Close() error
}

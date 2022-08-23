package redis_repo

import (
	"fmt"
	"github.com/go-redis/redis/v8"
)

type Config struct {
	Host     string
	Port     string
	Password string
}

func NewRedisClient(cfg Config) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", cfg.Host, cfg.Port),
		Password: cfg.Password,
		DB:       0,
	})

	return rdb
}

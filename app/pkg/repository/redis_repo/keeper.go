package redis_repo

import (
	"context"
	"errors"
	"sync"
	"time"

	"github.com/go-redis/redis/v8"
)


type RedisKeeper struct {
	rdb *redis.Client
	ctx context.Context
	mu  sync.Mutex
}

func NewRedisKeeper(rdb *redis.Client) *RedisKeeper {
	return &RedisKeeper{
		rdb: rdb,
		ctx: context.Background(),
		mu:  sync.Mutex{},
	}
}

func (rk *RedisKeeper) Get(key string) (string, error) {
	rk.mu.Lock()
	defer rk.mu.Unlock()

	val, err := rk.rdb.Get(rk.ctx, key).Result()
	if err == redis.Nil {
		return "", errors.New("message not found")
	}

	if err = rk.Clean(key); err != nil {
		return "", err
	}

	return val, err
}

func (rk *RedisKeeper) Set(key, message string, ttl int) error {
	return rk.rdb.Set(rk.ctx, key, message, time.Duration(ttl * 1000000000)).Err()
}

func (rk *RedisKeeper) Clean(key string) error {
	return rk.rdb.Del(rk.ctx, key).Err()
}

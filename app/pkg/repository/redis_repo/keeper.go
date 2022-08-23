package redis_repo

import (
	"context"
	"errors"
	"github.com/go-redis/redis/v8"
	"sync"
)

const TTL = 0

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
	val, err := rk.rdb.Get(rk.ctx, key).Result()
	if err == redis.Nil {
		rk.mu.Unlock()
		return "", errors.New("message not found")
	}

	if err = rk.Clean(key); err != nil {
		rk.mu.Unlock()
		return "", err
	}

	rk.mu.Unlock()
	return val, err
}

func (rk *RedisKeeper) Set(key, message string) error {
	return rk.rdb.Set(rk.ctx, key, message, TTL).Err()
}

func (rk *RedisKeeper) Clean(key string) error {
	return rk.rdb.Del(rk.ctx, key).Err()
}

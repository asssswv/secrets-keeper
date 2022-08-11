package redis_repo

import (
	"context"
	"errors"
	"github.com/go-redis/redis/v8"
)

const TTL = 0

type RedisKeeper struct {
	rdb *redis.Client
	ctx context.Context
}

func NewRedisKeeper(rdb *redis.Client) *RedisKeeper {
	return &RedisKeeper{
		rdb: rdb,
		ctx: context.Background(),
	}
}

func (rk *RedisKeeper) Get(key string) (string, error) {
	val, err := rk.rdb.Get(rk.ctx, key).Result()
	if err == redis.Nil {
		return "", errors.New("message not found")
	}

	return val, err
}

func (rk *RedisKeeper) Set(key, message string) error {
	return rk.rdb.Set(rk.ctx, key, message, TTL).Err()
}

func (rk *RedisKeeper) Clean(key string) error {
	return rk.rdb.Del(rk.ctx, key).Err()
}

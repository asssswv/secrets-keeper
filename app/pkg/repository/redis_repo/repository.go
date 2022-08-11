package redis_repo

import "github.com/go-redis/redis/v8"

type Keeper interface {
	Get(key string) (string, error)
	Set(key, message string) error
	Clean(key string) error
}

type Repository struct {
	Keeper
}

func NewRepository(rdb *redis.Client) *Repository {
	return &Repository{
		Keeper: NewRedisKeeper(rdb),
	}
}

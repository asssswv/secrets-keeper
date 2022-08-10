package redis_repo

import "github.com/go-redis/redis/v9"

type Keeper interface {
	Get(key string) (string, error)
	Set(key, message string) error
}

type Repository struct {
	Keeper
}

func NewRepository(rdb *redis.Client) *Repository {
	return &Repository{
		Keeper: NewRedisKeeper(rdb),
	}
}

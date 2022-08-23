package service

import (
	"github.com/stretchr/testify/assert"
	"secrets_keeper/app/pkg/repository/redis_repo"
	"testing"
)

// RaceTest
func TestKeeperService_Get(t *testing.T) {
	rdb := redis_repo.NewRedisClient(redis_repo.Config{
		Host:     "localhost",
		Port:     "6379",
		Password: "pass",
	})

	repos := redis_repo.NewRepository(rdb)
	services := NewService(repos)

	testMessage := "test"
	key, _ := services.UUIDKeyBuilder.Get()
	_ = services.Set(key, testMessage, 2)

	resultChannel := make(chan error, 2)

	go func(aKey string, c chan error) {
		_, err := services.Keeper.Get(aKey)
		c <- err
	}(key, resultChannel)

	go func(aKey string, c chan error) {
		_, err := services.Keeper.Get(aKey)
		c <- err
	}(key, resultChannel)

	firstErr := <-resultChannel
	secondErr := <-resultChannel

	// fmt.Println(firstErr, "===============", secondErr)

	assert.Equal(t, nil, firstErr)
	assert.NotEqual(t, nil, secondErr)

	assert.Equal(t, nil, firstErr)
	assert.Equal(t, "message not found", secondErr.Error())
}

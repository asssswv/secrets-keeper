package service

import (
	"secrets_keeper/app/pkg/repository/redis_repo"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
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

// ttl test
func TestKeeperService_Set(t *testing.T) {
	rdb := redis_repo.NewRedisClient(redis_repo.Config{
		Host:     "localhost",
		Port:     "6379",
		Password: "pass",
	})

	repos := redis_repo.NewRepository(rdb)
	services := NewService(repos)

	testTable := []struct {
		name    string
		message string
		key     string
		ttl     int
		wantErr bool
	}{
		{
			name:    "FAIL",
			message: "first",
			key:     "first_1",
			ttl:     2,
			wantErr: true,
		},
		{
			name:    "OK",
			message: "second",
			key:     "second_2",
			ttl:     100,
			wantErr: false,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			services.Keeper.Set(testCase.key, testCase.message, testCase.ttl)
			time.Sleep(3 * time.Second)

			message, err := services.Keeper.Get(testCase.key)
			if testCase.wantErr {
				assert.NotEqual(t, err, nil)
				assert.Equal(t, message, "")
			} else {
				// fmt.Println("===================", err)
				assert.Equal(t, err, nil)
				assert.NotEqual(t, message, nil)
			}
		})
	}
}

// validation test
func TestKeeperService_SetMessageLength(t *testing.T) {
	msg1 := ""
	msg2 := ""
	for i := 0; i < 1023; i++ {
		msg1 += "a"
		msg2 += "b"
	}

	msg1 += "a "

	rdb := redis_repo.NewRedisClient(redis_repo.Config{
		Host:     "localhost",
		Port:     "6379",
		Password: "pass",
	})

	repos := redis_repo.NewRepository(rdb)
	services := NewService(repos)

	testTable := []struct {
		name    string
		message string
		key     string
		ttl     int
		wantErr bool
	}{
		{
			name:    "FAIL",
			message: msg1,
			key:     "first_1",
			ttl:     2,
			wantErr: true,
		},
		{
			name:    "OK",
			message: msg2,
			key:     "second_2",
			ttl:     2,
			wantErr: false,
		},
	}

	for _, testCase := range testTable {
		err := services.Keeper.Set(testCase.key, testCase.message, testCase.ttl)

		if testCase.wantErr {
			assert.NotEqual(t, err, nil)
		} else {
			assert.Equal(t, err, nil)
		}
	}
}

func TestKeeperService_SetTTL(t *testing.T) {
	rdb := redis_repo.NewRedisClient(redis_repo.Config{
		Host:     "localhost",
		Port:     "6379",
		Password: "pass",
	})

	repos := redis_repo.NewRepository(rdb)
	services := NewService(repos)

	testTable := []struct {
		name    string
		message string
		key     string
		ttl     int
		wantErr bool
	}{
		{
			name:    "FAIL",
			message: "first",
			key:     "first_1",
			ttl:     100000,
			wantErr: true,
		},
		{
			name:    "OK",
			message: "second",
			key:     "second_2",
			ttl:     86300,
			wantErr: false,
		},
	}

	for _, testCase := range testTable {
		err := services.Keeper.Set(testCase.key, testCase.message, testCase.ttl)

		if testCase.wantErr {
			assert.NotEqual(t, err, nil)
		} else {
			assert.Equal(t, err, nil)
		}
	}
}

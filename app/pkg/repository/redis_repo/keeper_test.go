package redis_repo

import (
	"github.com/go-redis/redis/v8"
	"github.com/go-redis/redismock/v8"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRedisKeeper_Get(t *testing.T) {
	rdb, mock := redismock.NewClientMock()
	defer func(rdb *redis.Client) {
		err := rdb.Close()
		if err != nil {
			logrus.Fatal("failed to close rdb")
		}
	}(rdb)

	r := NewRedisKeeper(rdb)

	type mockBehavior func(key string)

	testTable := []struct {
		name         string
		mockBehavior mockBehavior
		key          string
		wantErr      bool
	}{
		{
			name: "FAIL",
			key:  "test",
			mockBehavior: func(key string) {
				mock.ExpectGet(key).RedisNil()
			},
			wantErr: true,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.mockBehavior(testCase.key)

			message, err := r.Get(testCase.key)
			if testCase.wantErr {
				assert.Equal(t, message, "")
				assert.NotEqual(t, err, redis.Nil)
			} else {
				assert.NotEqual(t, message, "")
				assert.Equal(t, err, redis.Nil)
			}
		})
	}
}

func TestRedisKeeper_Set(t *testing.T) {
	rdb, mock := redismock.NewClientMock()
	defer func(rdb *redis.Client) {
		err := rdb.Close()
		if err != nil {
			logrus.Fatal("failed to close rdb")
		}
	}(rdb)

	r := NewRedisKeeper(rdb)

	type args struct {
		key     string
		message string
	}

	type mockBehavior func(args args)

	testTable := []struct {
		name         string
		mockBehavior mockBehavior
		args         args
		wantErr      bool
	}{
		{
			name: "OK",
			args: args{
				key:     "test1",
				message: "hello",
			},
			mockBehavior: func(args args) {
				mock.ExpectSet(args.key, args.message, 0).RedisNil()
			},
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.mockBehavior(testCase.args)

			err := r.Set(testCase.args.key, testCase.args.message)
			if testCase.wantErr {
				assert.NotEqual(t, err, redis.Nil)
			} else {
				assert.Equal(t, err, redis.Nil)
			}
		})
	}
}

package service

import (
	"errors"
	"secrets_keeper/app/pkg/repository/redis_repo"
)

var MessageMaxLen = 1024
var MaxTTL = 86400

type KeeperService struct {
	repo redis_repo.Keeper
}

func NewKeeperService(repo redis_repo.Keeper) *KeeperService {
	return &KeeperService{repo: repo}
}

func (ks *KeeperService) Get(key string) (string, error) {
	message, err := ks.repo.Get(key)
	if err != nil {
		return "", err
	}

	return message, nil
}

func (ks *KeeperService) Set(key, message string, ttl int) error {
	if !validateMessageLength(message) {
		return errors.New("message length too long!")
	}

	if !validateMessageTTL(ttl) {
		return errors.New("ttl too long")
	}

	return ks.repo.Set(key, message, ttl)
}

func validateMessageLength(msg string) bool {
	return len(msg) < MessageMaxLen
}

func validateMessageTTL(ttl int) bool {
	return ttl < MaxTTL
}
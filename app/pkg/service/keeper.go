package service

import (
	"secrets_keeper/app/pkg/repository/redis_repo"
)

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

	if err = ks.repo.Clean(key); err != nil {
		return "", err
	}

	return message, nil
}

func (ks *KeeperService) Set(key, message string) error {
	return ks.repo.Set(key, message)
}

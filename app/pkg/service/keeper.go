package service

import "secrets_keeper/app/pkg/repository"

type KeeperService struct {
	repo repository.Keeper
}

func NewKeeperService(repo repository.Keeper) *KeeperService {
	return &KeeperService{repo: repo}
}

func (ks *KeeperService) Get(key string) (string, error) {
	return ks.repo.Get(key)
}
	
func (ks *KeeperService) Set(key, message string) error {
	return ks.repo.Set(key, message)
}
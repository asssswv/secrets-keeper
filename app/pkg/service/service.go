package service

import (
	"secrets_keeper/app/pkg/repository"
)

type Keeper interface {
	Get(key string) (string, error)
	Set(key, message string) error
}

type Service struct {
	Keeper
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Keeper: NewKeeperService(repos.Keeper),
	}
}
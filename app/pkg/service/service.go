package service

import (
	"secrets_keeper/app/pkg/repository/redis_repo"
)

//go:generate mockgen -source=service.go -destination=mocks/mock.go

type Keeper interface {
	Get(key string) (string, error)
	Set(key, message string) error
}

type UUIDKeyBuilder interface {
	Get() (string, error)
}

type Service struct {
	Keeper
	UUIDKeyBuilder
}

func NewService(repos *redis_repo.Repository) *Service {
	return &Service{
		Keeper:         NewKeeperService(repos.Keeper),
		UUIDKeyBuilder: NewUUIDKeyBuilderService(),
	}
}

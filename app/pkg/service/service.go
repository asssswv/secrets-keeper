package service

import (
	"secrets_keeper/app/pkg/repository"
)

type Keeper interface {
	Get(key string) (string, error)
	Set(key, message string) error
}

type KeyBuilder interface {
	Get() (string, error)
}

type UUIDKeyBuilder interface {
	Get() (string, error)
}

type Service struct {
	Keeper
	KeyBuilder
	UUIDKeyBuilder
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Keeper:         NewKeeperService(repos.Keeper),
		KeyBuilder:     NewKeyBuilderService(),
		UUIDKeyBuilder: NewUUIDKeyBuilderService(),
	}
}

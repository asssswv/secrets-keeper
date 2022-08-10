package service

import "github.com/google/uuid"

type UUIDKeyBuilderService struct {
}

func NewUUIDKeyBuilderService() *UUIDKeyBuilderService {
	return &UUIDKeyBuilderService{}
}

func (_ *UUIDKeyBuilderService) Get() (string, error) {
	key, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}

	return key.String(), nil
}

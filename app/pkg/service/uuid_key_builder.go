package service

import "github.com/google/uuid"

type UUIDKeyBuilderService struct {
}

func NewUUIDKeyBuilderService() *UUIDKeyBuilderService {
	return &UUIDKeyBuilderService{}
}

func (ukbs *UUIDKeyBuilderService) Get() (string, error) {
	uuid, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}

	return uuid.String(), nil
}

package service

type KeyBuilderService struct {
}

func NewKeyBuilderService() *KeyBuilderService {
	return &KeyBuilderService{}
}

func (kbs *KeyBuilderService) Get() string {
	return "test"
}

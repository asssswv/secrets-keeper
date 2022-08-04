package repository

type Keeper interface {
	Get(key string) (string, error)
	Set(key, message string) error
	Clean(key string) error
}

type Repository struct {
	mem map[string]string
	Keeper
}

func NewRepository(mem map[string]string) *Repository {
	return &Repository{
		mem:    mem,
		Keeper: NewKeeperLocalMem(mem),
	}
}

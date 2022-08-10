package local_mem

type Keeper interface {
	Get(key string) (string, error)
	Set(key, message string) error
	Clean(key string) error
}

type LocalMemRepository struct {
	mem map[string]string
	Keeper
}

func NewRepository(mem map[string]string) *LocalMemRepository {
	return &LocalMemRepository{
		mem:    mem,
		Keeper: NewKeeperLocalMem(mem),
	}
}

package local_mem

import (
	"errors"
)

type KeeperLocalMem struct {
	mem map[string]string
}

func NewKeeperLocalMem(mem map[string]string) *KeeperLocalMem {
	return &KeeperLocalMem{
		mem: mem,
	}
}

func (k *KeeperLocalMem) Get(key string) (string, error) {
	value, ok := k.mem[key]
	if !ok {
		return "", errors.New("message not found")
	}

	return value, nil
}

func (k *KeeperLocalMem) Set(key, message string) error {
	k.mem[key] = message
	return nil
}

func (k *KeeperLocalMem) Clean(key string) error {
	delete(k.mem, key)
	return nil
}

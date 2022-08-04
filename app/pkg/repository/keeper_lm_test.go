package repository

import "testing"

func TestKeeperSet(t *testing.T) {
	keeper := KeeperLocalMem{mem: make(map[string]string)}
	key := "pupa"
	value := "lupa"

	_ = keeper.Set(key, value)

	if keeper.mem[key] != value {
		t.Error("bad memory value")
	}
}

func TestKeeperGet(t *testing.T) {
	keeper := KeeperLocalMem{mem: make(map[string]string)}
	key := "pupa"
	expected_value := "lupa"

	keeper.mem[key] = expected_value

	real_value, _ := keeper.Get(key)

	if real_value != expected_value {
		t.Error("bad  real value")
	}
}

func TestKeeperClear(t *testing.T) {
	keeper := KeeperLocalMem{mem: make(map[string]string)}
	key := "pupa"
	expected_value := "lupa"

	keeper.mem[key] = expected_value
	keeper.Clean(key)

	if _, ok := keeper.mem[key]; ok {
		t.Error("clean dosn't work") 
	}
}

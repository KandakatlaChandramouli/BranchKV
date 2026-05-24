package tests

import (
	"branchkv-core/internal/arena"
	"testing"
)

func TestArenaAllocate(t *testing.T) {

	a := arena.NewArena()

	page := a.Allocate(128)

	if len(page) == 0 {
		t.Fatal("allocation failed")
	}
}

func TestArenaFree(t *testing.T) {

	a := arena.NewArena()

	page := a.Allocate(128)

	a.Free(page)
}

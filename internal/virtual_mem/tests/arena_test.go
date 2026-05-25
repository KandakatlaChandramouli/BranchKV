package tests

import (
	"branchkv-core/internal/virtual_mem"
	"testing"
)

func TestArenaAllocate(
	t *testing.T,
) {

	arena := virtual_mem.NewArena(
		16,
		128,
	)

	frame, ok := arena.Allocate()

	if !ok {
		t.Fatal("allocation failed")
	}

	if frame.ID > 15 {
		t.Fatal("invalid frame id")
	}
}

package tests

import (
	"branchkv-core/internal/memoryarena"
	"testing"
)

func TestArena(
	t *testing.T,
) {

	a := memoryarena.NewArena(
		1024,
	)

	v := a.Allocate(128)

	if len(v) != 128 {
		t.Fatal("arena failed")
	}
}

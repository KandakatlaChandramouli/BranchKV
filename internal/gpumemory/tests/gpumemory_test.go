package tests

import (
	"branchkv-core/internal/gpumemory"
	"testing"
)

func TestGPUMemory(
	t *testing.T,
) {

	pool := gpumemory.NewPool()

	pool.Allocate(1024)

	if pool.Usage() != 1024 {
		t.Fatal("gpu memory failed")
	}
}

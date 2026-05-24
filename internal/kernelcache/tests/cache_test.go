package tests

import (
	"branchkv-core/internal/kernelcache"
	"testing"
)

func TestKernelCache(
	t *testing.T,
) {

	c := kernelcache.NewCache()

	c.Store("matmul")

	if c.Size() != 1 {
		t.Fatal("cache failed")
	}
}

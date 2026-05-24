package tests

import (
	"branchkv-core/internal/branchcache"
	"branchkv-core/internal/types"
	"testing"
)

func TestBranchCache(
	t *testing.T,
) {

	c := branchcache.NewBranchCache()

	seq := []types.TokenID{
		1, 2, 3,
	}

	c.Store(1, seq)

	v, ok := c.Load(1)

	if !ok {
		t.Fatal("cache miss")
	}

	if len(v) != 3 {
		t.Fatal("invalid seq")
	}
}

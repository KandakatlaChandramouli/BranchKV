package tests

import (
	"branchkv-core/internal/cache"
	"branchkv-core/internal/types"
	"testing"
)

func TestTokenCache(
	t *testing.T,
) {

	c := cache.NewTokenCache()

	vec := []float32{
		1, 2, 3,
	}

	c.Store(
		types.TokenID(1),
		vec,
	)

	v, ok := c.Load(1)

	if !ok {
		t.Fatal("cache miss")
	}

	if len(v) != 3 {
		t.Fatal("invalid vector")
	}
}

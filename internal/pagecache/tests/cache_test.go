package tests

import (
	"branchkv-core/internal/pagecache"
	"testing"
)

func TestPageCache(
	t *testing.T,
) {

	cache := pagecache.NewPageCache()

	cache.Store(
		&pagecache.CachedPage{
			ID: 1,
			Data: []byte{
				1, 2, 3,
			},
		},
	)

	_, ok := cache.Load(1)

	if !ok {
		t.Fatal("cache miss")
	}
}

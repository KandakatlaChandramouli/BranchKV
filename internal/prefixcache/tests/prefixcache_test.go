package tests

import (
	"branchkv-core/internal/prefixcache"
	"testing"
)

func TestPrefixCache(
	t *testing.T,
) {

	c := prefixcache.NewCache()

	c.Store(
		"hello",
		[]float32{
			1, 2, 3,
		},
	)

	if c.Size() != 1 {
		t.Fatal("cache failed")
	}
}

package mega

import (
	"branchkv-core/internal/vectorcache"
	"testing"
)

func BenchmarkVectorCache(
	b *testing.B,
) {

	cache := vectorcache.NewRuntime()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		cache.Add()
	}
}

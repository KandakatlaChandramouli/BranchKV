package runtime

import (
	"branchkv-core/internal/cache"
	"branchkv-core/internal/metrics"
	"branchkv-core/internal/prefetch"
	"branchkv-core/internal/types"
	"testing"
)

func BenchmarkRuntimePipeline(
	b *testing.B,
) {

	cacheEngine := cache.NewTokenCache()

	metricsEngine := metrics.NewMetrics()

	prefetcher := prefetch.NewPrefetcher(4)

	seq := []types.TokenID{
		1, 2, 3, 4, 5, 6, 7, 8,
	}

	cacheEngine.Store(
		1,
		[]float32{
			1, 2, 3,
		},
	)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		_, _ = cacheEngine.Load(1)

		_ = prefetcher.Predict(seq)

		metricsEngine.RecordTrieSearch()
	}
}

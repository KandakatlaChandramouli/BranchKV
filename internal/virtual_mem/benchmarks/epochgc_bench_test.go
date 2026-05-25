package benchmarks

import (
	"branchkv-core/internal/virtual_mem/epochgc"
	"testing"
)

func BenchmarkEpochAdvance(
	b *testing.B,
) {

	gc := epochgc.NewCollector()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		gc.Advance()
	}
}

package benchmarks

import (
	"branchkv-core/internal/virtual_mem/fixeddag"
	"branchkv-core/internal/virtual_mem/specarena"
	"testing"
)

func BenchmarkFixedDAG(
	b *testing.B,
) {

	total := uint64(
		(b.N * (fixeddag.MaxChildren + 1)) + 1,
	)

	arena := specarena.NewArena(
		total,
	)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		root := arena.Allocate()

		for j := 0; j < fixeddag.MaxChildren; j++ {

			child := arena.Allocate()

			fixeddag.Link(
				root,
				child,
			)
		}
	}
}

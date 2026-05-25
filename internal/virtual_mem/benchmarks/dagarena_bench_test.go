package benchmarks

import (
	"branchkv-core/internal/virtual_mem/dagnodearena"
	"branchkv-core/internal/virtual_mem/specbranchgraph"
	"testing"
)

func BenchmarkArenaDAG(
	b *testing.B,
) {

	arena := dagnodearena.NewArena(
		uint64(
			b.N + 1,
		),
	)

	root := arena.Allocate()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		child := arena.Allocate()

		specbranchgraph.Link(
			root,
			child,
		)
	}
}

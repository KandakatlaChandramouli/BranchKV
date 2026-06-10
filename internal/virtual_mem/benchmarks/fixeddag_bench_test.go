package benchmarks

import (
	"branchkv-core/internal/virtual_mem/fixeddag"
	"branchkv-core/internal/virtual_mem/specarena"
	"testing"
)

func BenchmarkFixedDAG(
	b *testing.B,
) {

	const arenaSize = 100000

	arena := specarena.NewArena(
		arenaSize,
	)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		root := arena.Allocate()

		if root == nil {
			arena = specarena.NewArena(
				arenaSize,
			)
			root = arena.Allocate()
		}

		for j := 0; j < fixeddag.MaxChildren; j++ {

			child := arena.Allocate()

			if child == nil {

				arena = specarena.NewArena(
					arenaSize,
				)

				child = arena.Allocate()
			}

			fixeddag.Link(
				root,
				child,
			)
		}
	}
}

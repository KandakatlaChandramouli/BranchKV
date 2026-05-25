package benchmarks

import (
	"branchkv-core/internal/virtual_mem/specdag"
	"testing"
)

func BenchmarkSpecDAG(
	b *testing.B,
) {

	root := specdag.NewNode(
		1,
	)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		child := specdag.NewNode(
			uint64(i),
		)

		specdag.Link(
			root,
			child,
		)
	}
}

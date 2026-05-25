package benchmarks

import (
	"branchkv-core/internal/virtual_mem/pagetable/l1"
	"branchkv-core/internal/virtual_mem/snapshotspace"
	"branchkv-core/internal/virtual_mem/specwalk"
	"branchkv-core/internal/virtual_mem/tlb"
	"testing"
)

func BenchmarkSpecWalker(
	b *testing.B,
) {

	root := l1.NewTable()

	space := snapshotspace.NewSpace(
		1,
		root,
	)

	cache := tlb.NewCache()

	walker := specwalk.NewWalker(
		space,
		cache,
	)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		_ = walker.Inner()
	}
}

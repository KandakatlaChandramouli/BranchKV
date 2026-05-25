package benchmarks

import (
	"branchkv-core/internal/virtual_mem"
	"branchkv-core/internal/virtual_mem/pagetable/l1"
	"branchkv-core/internal/virtual_mem/pagewalker"
	"branchkv-core/internal/virtual_mem/tlb"
	"testing"
)

func BenchmarkPageWalker(
	b *testing.B,
) {

	root := l1.NewTable()

	cache := tlb.NewCache()

	walker := pagewalker.NewWalker(
		root,
		cache,
	)

	l2Table := root.Ensure(
		1,
	)

	frame := virtual_mem.NewPhysicalFrame(
		1,
		256,
	)

	l2Table.Map(
		2,
		frame,
	)

	virtual := uint64(
		(1 << 22) |
			(2 << 12),
	)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		walker.Translate(
			virtual,
		)
	}
}

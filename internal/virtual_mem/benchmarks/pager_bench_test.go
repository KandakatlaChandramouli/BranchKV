package benchmarks

import (
	"branchkv-core/internal/virtual_mem"
	"branchkv-core/internal/virtual_mem/eviction"
	"branchkv-core/internal/virtual_mem/pager"
	"branchkv-core/internal/virtual_mem/swap"
	"testing"
)

func BenchmarkPageFault(
	b *testing.B,
) {

	swapper := swap.NewManager()

	evictor := eviction.NewEvictor(
		swapper,
	)

	p := pager.NewPager(
		swapper,
	)

	frame := virtual_mem.NewPhysicalFrame(
		1,
		128,
	)

	frame.MarkDirty()

	evictor.Evict(
		frame,
	)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		p.Fault(
			frame,
		)
	}
}

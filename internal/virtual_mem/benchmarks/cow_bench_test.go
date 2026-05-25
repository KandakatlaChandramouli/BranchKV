package benchmarks

import (
	"branchkv-core/internal/virtual_mem"
	"branchkv-core/internal/virtual_mem/cow"
	"testing"
)

func BenchmarkCopyOnWrite(
	b *testing.B,
) {

	allocator := virtual_mem.NewFrameAllocator(
		4096,
		256,
	)

	table := virtual_mem.NewPageTable()

	frame, _ := allocator.Alloc()

	table.Map(
		1,
		frame,
	)

	engine := cow.NewCopyOnWrite(
		allocator,
		table,
	)

	engine.Fork(
		1,
		2,
	)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		engine.Write(
			2,
			0,
			float32(i),
		)
	}
}

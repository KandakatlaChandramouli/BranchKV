package benchmarks

import (
	"branchkv-core/internal/virtual_mem"
	"testing"
)

func BenchmarkArenaAllocate(
	b *testing.B,
) {

	arena := virtual_mem.NewArena(
		4096,
		256,
	)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		frame, ok := arena.Allocate()

		if !ok {
			b.Fatal("allocation failed")
		}

		arena.Free(frame)
	}
}

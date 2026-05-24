package benchmarks

import (
	"branchkv-core/internal/arena"
	"testing"
)

func BenchmarkArenaAllocate(b *testing.B) {

	a := arena.NewArena()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		page := a.Allocate(256)

		a.Free(page)
	}
}

package benchmarks

import (
	"branchkv-core/internal/virtual_mem"
	"testing"
)

func BenchmarkFork(b *testing.B) {
	mmu := virtual_mem.NewMMU()

	mmu.Allocate(1, 4096)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = mmu.Fork(1, uint64(i+2))
	}
}

package mmu

import (
	"branchkv-core/internal/virtual_mem"
	"testing"
)

func BenchmarkMMUWrite(
	b *testing.B,
) {

	mmu := virtual_mem.NewMMU()

	_ = mmu.AllocatePage(
		1,
	)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		_ = mmu.Write(
			1,
			0,
			float32(i),
		)
	}
}

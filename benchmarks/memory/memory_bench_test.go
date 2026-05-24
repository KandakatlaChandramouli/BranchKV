package memory

import (
	"branchkv-core/internal/slaballocator"
	"testing"
)

func BenchmarkMemoryAllocator(
	b *testing.B,
) {

	a := slaballocator.NewAllocator()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		_ = a.Allocate(64)
	}
}

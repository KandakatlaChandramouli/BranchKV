package benchmarks

import (
	"branchkv-core/internal/unsafe_mem"
	"testing"
)

func BenchmarkUnsafeWriteRead(
	b *testing.B,
) {

	page := unsafe_mem.NewRawPage(4096)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		page.WriteFloat32(0, 99)

		_ = page.ReadFloat32(0)
	}
}

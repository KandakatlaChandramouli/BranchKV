package benchmarks

import (
	"branchkv-core/internal/virtual_mem"
	"branchkv-core/internal/virtual_mem/cowfault"
	"testing"
)

func BenchmarkCoWFault(
	b *testing.B,
) {

	frame := virtual_mem.NewPhysicalFrame(
		1,
		256,
	)

	frame.RefCount.Store(
		2,
	)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		_ = cowfault.Handle(
			frame,
		)
	}
}

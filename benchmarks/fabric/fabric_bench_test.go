package fabric

import (
	"branchkv-core/internal/fabric"
	"testing"
)

func BenchmarkFabric(
	b *testing.B,
) {

	f := fabric.NewFabric()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		f.AddNode(
			uint64(i),
		)
	}
}

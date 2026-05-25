package distributed_storage

import (
	"branchkv-core/internal/vectorindex"
	"testing"
)

func BenchmarkVectorInsert(
	b *testing.B,
) {

	idx := vectorindex.NewIndex()

	vec := []float32{
		1, 2, 3, 4,
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		idx.Insert(
			uint64(i),
			vec,
		)
	}
}

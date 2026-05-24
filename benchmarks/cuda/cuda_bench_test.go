package cuda

import (
	"branchkv-core/internal/allreduce"
	"testing"
)

func BenchmarkCUDAFabric(
	b *testing.B,
) {

	input := [][]float32{
		{1, 2, 3},
		{4, 5, 6},
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		_ = allreduce.Sum(input)
	}
}

package tensor

import (
	"branchkv-core/internal/matmul"
	"testing"
)

func BenchmarkTensorRuntime(
	b *testing.B,
) {

	a := [][]float32{
		{1, 2},
		{3, 4},
	}

	c := [][]float32{
		{5, 6},
		{7, 8},
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		_ = matmul.Multiply(
			a,
			c,
		)
	}
}

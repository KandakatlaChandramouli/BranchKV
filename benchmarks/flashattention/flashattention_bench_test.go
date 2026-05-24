package flashattention

import (
	"branchkv-core/internal/flashattention"
	"testing"
)

func BenchmarkFlashAttention(
	b *testing.B,
) {

	engine := flashattention.NewFlashAttention()

	q := []float64{
		1, 2, 3, 4,
	}

	k := []float64{
		5, 6, 7, 8,
	}

	v := []float64{
		9, 10, 11, 12,
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		_ = engine.Compute(
			q,
			k,
			v,
		)
	}
}

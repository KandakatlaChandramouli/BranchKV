package transformer

import (
	"branchkv-core/internal/embedding"
	"branchkv-core/internal/softmax"
	"branchkv-core/internal/transformer"
	"testing"
)

func BenchmarkTransformerRuntime(
	b *testing.B,
) {

	t := transformer.NewTransformer()

	e := embedding.NewEmbedding(
		1000,
		64,
	)

	input := []float64{
		1, 2, 3, 4,
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		t.Forward()

		_ = e.Encode(1)

		_ = softmax.Compute(input)
	}
}

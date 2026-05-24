package pipeline

import (
	"branchkv-core/internal/pipeline"
	"branchkv-core/internal/types"
	"testing"
)

func BenchmarkPipeline(
	b *testing.B,
) {

	p := pipeline.NewPipeline()

	seq := []types.TokenID{
		1, 2, 3, 4, 5, 6, 7, 8,
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = p.Process(seq)
	}
}

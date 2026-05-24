package speculative

import (
	"branchkv-core/internal/speculative"
	"branchkv-core/internal/types"
	"testing"
)

func BenchmarkSpeculativeFork(
	b *testing.B,
) {

	engine := speculative.NewSpeculativeEngine()

	seq := []types.TokenID{
		1, 2, 3, 4, 5, 6, 7, 8,
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = engine.ForkSequence(seq)
	}
}

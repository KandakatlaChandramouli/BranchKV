package tests

import (
	"branchkv-core/internal/flashattention"
	"testing"
)

func TestFlashAttention(
	t *testing.T,
) {

	engine := flashattention.NewFlashAttention()

	out := engine.Compute(
		[]float64{1, 2},
		[]float64{3, 4},
		[]float64{5, 6},
	)

	if len(out) != 2 {
		t.Fatal("flash attention failed")
	}
}

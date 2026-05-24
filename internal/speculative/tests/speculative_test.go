package tests

import (
	"branchkv-core/internal/speculative"
	"branchkv-core/internal/types"
	"testing"
)

func TestForkSequence(
	t *testing.T,
) {

	engine := speculative.NewSpeculativeEngine()

	seq := []types.TokenID{
		1, 2, 3,
	}

	out := engine.ForkSequence(seq)

	if len(out) != 3 {
		t.Fatal("fork failed")
	}
}

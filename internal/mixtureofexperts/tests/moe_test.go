package tests

import (
	"branchkv-core/internal/mixtureofexperts"
	"testing"
)

func TestMoE(
	t *testing.T,
) {

	moe := mixtureofexperts.NewMoE()

	moe.AddExpert(1)

	if moe.Size() != 1 {
		t.Fatal("moe failed")
	}
}

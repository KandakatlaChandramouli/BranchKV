package tests

import (
	"branchkv-core/internal/simd"
	"branchkv-core/internal/types"
	"testing"
)

func TestCompareTokens(
	t *testing.T,
) {

	a := []types.TokenID{
		1, 2, 3, 4,
	}

	b := []types.TokenID{
		1, 2, 3, 4,
	}

	ok := simd.CompareTokens(a, b)

	if !ok {
		t.Fatal("token compare failed")
	}
}

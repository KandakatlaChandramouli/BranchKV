package tests

import (
	"branchkv-core/internal/matmul"
	"testing"
)

func TestMatmul(
	t *testing.T,
) {

	out := matmul.Multiply(
		[][]float32{
			{1, 2},
		},
		[][]float32{
			{3},
			{4},
		},
	)

	if out[0][0] != 11 {
		t.Fatal("matmul failed")
	}
}

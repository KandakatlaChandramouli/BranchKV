package tests

import (
	"branchkv-core/internal/tensor"
	"testing"
)

func TestTensor(
	t *testing.T,
) {

	tensor := tensor.NewTensor(
		[]int{2, 2},
		[]float32{
			1, 2, 3, 4,
		},
	)

	if tensor.Size() != 4 {
		t.Fatal("tensor failed")
	}
}

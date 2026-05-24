package tests

import (
	"branchkv-core/internal/softmax"
	"testing"
)

func TestSoftmax(
	t *testing.T,
) {

	out := softmax.Compute(
		[]float64{
			1, 2, 3,
		},
	)

	if len(out) != 3 {
		t.Fatal("softmax failed")
	}
}

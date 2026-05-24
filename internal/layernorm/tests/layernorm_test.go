package tests

import (
	"branchkv-core/internal/layernorm"
	"testing"
)

func TestLayerNorm(
	t *testing.T,
) {

	out := layernorm.Normalize(
		[]float32{
			1, 2, 3,
		},
	)

	if len(out) != 3 {
		t.Fatal("layernorm failed")
	}
}

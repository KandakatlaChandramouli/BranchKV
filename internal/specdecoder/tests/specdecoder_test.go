package tests

import (
	"branchkv-core/internal/specdecoder"
	"branchkv-core/internal/types"
	"testing"
)

func TestSpecDecoder(
	t *testing.T,
) {

	s := specdecoder.NewSpecDecoder()

	out := s.Predict(
		[]types.TokenID{
			1, 2, 3,
		},
	)

	if len(out) != 4 {
		t.Fatal("spec decode failed")
	}
}

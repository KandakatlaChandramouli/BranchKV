package tests

import (
	"branchkv-core/internal/beamsearch"
	"branchkv-core/internal/types"
	"testing"
)

func TestBeamSearch(
	t *testing.T,
) {

	b := beamsearch.NewBeamSearch(4)

	out := b.Expand(
		[]types.TokenID{
			1, 2,
		},
	)

	if len(out) != 4 {
		t.Fatal("beam failed")
	}
}

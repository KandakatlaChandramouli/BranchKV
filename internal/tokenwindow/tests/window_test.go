package tests

import (
	"branchkv-core/internal/tokenwindow"
	"branchkv-core/internal/types"
	"testing"
)

func TestWindow(
	t *testing.T,
) {

	w := tokenwindow.NewWindow(2)

	out := w.Slice(
		[]types.TokenID{
			1, 2, 3, 4,
		},
	)

	if len(out) != 2 {
		t.Fatal("window failed")
	}
}

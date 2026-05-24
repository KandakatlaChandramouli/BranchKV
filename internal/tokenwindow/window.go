package tokenwindow

import (
	"branchkv-core/internal/types"
)

type Window struct {
	Size int
}

func NewWindow(
	size int,
) *Window {

	return &Window{
		Size: size,
	}
}

func (w *Window) Slice(
	seq []types.TokenID,
) []types.TokenID {

	if len(seq) <= w.Size {
		return seq
	}

	return seq[len(seq)-w.Size:]
}

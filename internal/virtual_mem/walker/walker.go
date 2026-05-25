package walker

import (
	"branchkv-core/internal/virtual_mem/address"
	"branchkv-core/internal/virtual_mem/pagetable/l1"
	"branchkv-core/internal/virtual_mem/translation"
)

type Walker struct {
	root *l1.Table
}

func NewWalker(
	root *l1.Table,
) *Walker {

	return &Walker{
		root: root,
	}
}

func (w *Walker) Translate(
	virtual uint64,
) translation.Result {

	l1Index, l2Index, offset := address.Split(
		virtual,
	)

	l2Table, ok := w.root.Resolve(
		l1Index,
	)

	if !ok {
		return translation.Result{}
	}

	frame, ok := l2Table.Resolve(
		l2Index,
	)

	if !ok {
		return translation.Result{}
	}

	return translation.Result{
		Frame:  frame,
		Offset: offset,
		Hit:    true,
	}
}

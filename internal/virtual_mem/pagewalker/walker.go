package pagewalker

import (
	"branchkv-core/internal/virtual_mem/address"
	"branchkv-core/internal/virtual_mem/pagetable/l1"
	"branchkv-core/internal/virtual_mem/tlb"
	"branchkv-core/internal/virtual_mem/translation"
)

type Walker struct {
	root *l1.Table
	tlb  *tlb.Cache
}

func NewWalker(
	root *l1.Table,
	cache *tlb.Cache,
) *Walker {

	return &Walker{
		root: root,
		tlb:  cache,
	}
}

func (w *Walker) Translate(
	virtual uint64,
) translation.Result {

	frame, ok := w.tlb.Lookup(
		virtual,
	)

	if ok {

		_, _, offset := address.Split(
			virtual,
		)

		return translation.Result{
			Frame:  frame,
			Offset: offset,
			Hit:    true,
		}
	}

	l1Index, l2Index, offset := address.Split(
		virtual,
	)

	l2Table, ok := w.root.Resolve(
		l1Index,
	)

	if !ok {
		return translation.Result{}
	}

	frame, ok = l2Table.Resolve(
		l2Index,
	)

	if !ok {
		return translation.Result{}
	}

	w.tlb.Insert(
		virtual,
		frame,
	)

	return translation.Result{
		Frame:  frame,
		Offset: offset,
		Hit:    true,
	}
}

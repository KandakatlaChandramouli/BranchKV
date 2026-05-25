package tests

import (
	"branchkv-core/internal/virtual_mem"
	"branchkv-core/internal/virtual_mem/pagetable/l1"
	"branchkv-core/internal/virtual_mem/pagewalker"
	"branchkv-core/internal/virtual_mem/tlb"
	"testing"
)

func TestPageWalker(
	t *testing.T,
) {

	root := l1.NewTable()

	cache := tlb.NewCache()

	walker := pagewalker.NewWalker(
		root,
		cache,
	)

	l2Table := root.Ensure(
		1,
	)

	frame := virtual_mem.NewPhysicalFrame(
		99,
		128,
	)

	l2Table.Map(
		2,
		frame,
	)

	virtual := uint64(
		(1 << 22) |
			(2 << 12),
	)

	result := walker.Translate(
		virtual,
	)

	if !result.Hit {
		t.Fatal("translation failed")
	}

	if result.Frame.ID != 99 {
		t.Fatal("wrong frame")
	}
}

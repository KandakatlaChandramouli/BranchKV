package tests

import (
	"branchkv-core/internal/virtual_mem"
	"branchkv-core/internal/virtual_mem/cow"
	"testing"
)

func TestCopyOnWrite(
	t *testing.T,
) {

	allocator := virtual_mem.NewFrameAllocator(
		32,
		128,
	)

	table := virtual_mem.NewPageTable()

	frame, ok := allocator.Alloc()

	if !ok {
		t.Fatal("alloc failed")
	}

	table.Map(
		1,
		frame,
	)

	engine := cow.NewCopyOnWrite(
		allocator,
		table,
	)

	ok = engine.Fork(
		1,
		2,
	)

	if !ok {
		t.Fatal("fork failed")
	}

	ok = engine.Write(
		2,
		0,
		42,
	)

	if !ok {
		t.Fatal("cow write failed")
	}

	parent, _ := table.Resolve(1)
	child, _ := table.Resolve(2)

	if parent == child {
		t.Fatal("copy-on-write failed")
	}
}

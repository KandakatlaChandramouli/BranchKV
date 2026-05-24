package tests

import (
	"branchkv-core/internal/virtual_mem"
	"testing"
)

func TestMMUAllocate(
	t *testing.T,
) {

	mmu := virtual_mem.NewMMU()

	err := mmu.AllocatePage(1)

	if err != nil {
		t.Fatal(err)
	}
}

func TestMMUFork(
	t *testing.T,
) {

	mmu := virtual_mem.NewMMU()

	err := mmu.AllocatePage(1)

	if err != nil {
		t.Fatal(err)
	}

	err = mmu.Fork(
		1,
		2,
	)

	if err != nil {
		t.Fatal(err)
	}

	if mmu.RefCount(1) != 2 {
		t.Fatal("bad refcount")
	}
}

func TestMMUCopyOnWrite(
	t *testing.T,
) {

	mmu := virtual_mem.NewMMU()

	err := mmu.AllocatePage(1)

	if err != nil {
		t.Fatal(err)
	}

	err = mmu.Write(
		1,
		0,
		10,
	)

	if err != nil {
		t.Fatal(err)
	}

	err = mmu.Fork(
		1,
		2,
	)

	if err != nil {
		t.Fatal(err)
	}

	err = mmu.Write(
		2,
		0,
		99,
	)

	if err != nil {
		t.Fatal(err)
	}

	parent, _ := mmu.Read(
		1,
		0,
	)

	child, _ := mmu.Read(
		2,
		0,
	)

	if parent != 10 {
		t.Fatal("parent corrupted")
	}

	if child != 99 {
		t.Fatal("cow failed")
	}
}

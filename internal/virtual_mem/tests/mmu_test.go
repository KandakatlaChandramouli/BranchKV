package tests

import (
	"branchkv-core/internal/virtual_mem"
	"testing"
)

func TestForkSharesMemory(t *testing.T) {
	mmu := virtual_mem.NewMMU()

	original := mmu.Allocate(1, 8)

	err := mmu.Fork(1, 2)
	if err != nil {
		t.Fatal(err)
	}

	forked, _ := mmu.Resolve(2)

	if original.Page != forked.Page {
		t.Fatal("pages not shared")
	}
}

func TestCopyOnWrite(t *testing.T) {
	mmu := virtual_mem.NewMMU()

	mmu.Allocate(1, 8)

	mmu.Write(1, 0, 10)

	mmu.Fork(1, 2)

	before, _ := mmu.Resolve(1)
	shared := before.Page

	mmu.Write(2, 0, 99)

	after, _ := mmu.Resolve(2)

	if shared == after.Page {
		t.Fatal("copy-on-write failed")
	}
}

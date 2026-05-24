package tests

import (
	"branchkv-core/internal/tree"
	"branchkv-core/internal/types"
	"branchkv-core/internal/virtual_mem"
	"testing"
)

func TestInsertAndSearch(t *testing.T) {
	rt := tree.NewRadixTree()

	mmu := virtual_mem.NewMMU()

	desc := mmu.Allocate(1, 8)

	sequence := []types.TokenID{1, 2, 3, 4}

	rt.Insert(sequence, desc)

	node, found := rt.Search(sequence)

	if !found {
		t.Fatal("sequence not found")
	}

	if node.Descriptor != desc {
		t.Fatal("descriptor mismatch")
	}
}

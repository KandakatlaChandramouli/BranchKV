package tests

import (
	"branchkv-core/internal/tree"
	"branchkv-core/internal/types"
	"branchkv-core/internal/virtual_mem"
	"testing"
)

func TestRadixInsert(
	t *testing.T,
) {

	rt := tree.NewRadixTree()

	mmu := virtual_mem.NewMMU()

	desc, err := mmu.AllocateDescriptor(
		1,
		128,
	)

	if err != nil {
		t.Fatal(err)
	}

	sequence := []types.TokenID{
		1, 2, 3,
	}

	rt.Insert(
		sequence,
		desc,
	)

	_, ok := rt.Search(
		sequence,
	)

	if !ok {
		t.Fatal("search failed")
	}
}

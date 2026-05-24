package main

import (
	"branchkv-core/internal/tree"
	"branchkv-core/internal/types"
	"branchkv-core/internal/virtual_mem"
	"fmt"
)

func main() {

	mmu := virtual_mem.NewMMU()

	desc, err := mmu.AllocateDescriptor(
		1,
		256,
	)

	if err != nil {
		panic(err)
	}

	rt := tree.NewRadixTree()

	sequence := []types.TokenID{
		1, 2, 3, 4,
	}

	rt.Insert(
		sequence,
		desc,
	)

	found, ok := rt.Search(
		sequence,
	)

	if !ok {
		panic("sequence missing")
	}

	parent, ok := mmu.Resolve(
		found.VirtualPageID,
	)

	if !ok {
		panic("resolve failed")
	}

	fmt.Println(
		"page:",
		parent.ID,
	)

	fmt.Println(
		"refcount:",
		parent.RefCount.Load(),
	)
}

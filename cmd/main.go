package main

import (
	"branchkv-core/internal/tree"
	"branchkv-core/internal/types"
	"branchkv-core/internal/virtual_mem"
	"fmt"
)

func main() {

	mmu := virtual_mem.NewMMU()

	rt := tree.NewRadixTree()

	desc := mmu.Allocate(1, 16)

	sequence := []types.TokenID{
		101,
		102,
		103,
		104,
	}

	node := rt.Insert(sequence, desc)

	branchA := tree.NewBranch(
		1,
		sequence,
		node,
	)

	branchB, err := tree.ForkBranch(
		mmu,
		branchA,
		2,
	)

	if err != nil {
		panic(err)
	}

	err = mmu.Write(2, 0, 999)

	if err != nil {
		panic(err)
	}

	fmt.Println("Branch A ID:", branchA.ID)
	fmt.Println("Branch B ID:", branchB.ID)

	aDesc, _ := mmu.Resolve(1)
	bDesc, _ := mmu.Resolve(2)

	fmt.Println("A RefCount:", aDesc.Page.Count())
	fmt.Println("B RefCount:", bDesc.Page.Count())

	fmt.Println("A Page:", aDesc.Page.ID)
	fmt.Println("B Page:", bDesc.Page.ID)
}

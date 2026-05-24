package main

import (
	"branchkv-core/internal/virtual_mem"
	"fmt"
)

func main() {

	mmu := virtual_mem.NewMMU()

	err := mmu.AllocatePage(
		1,
	)

	if err != nil {
		panic(err)
	}

	err = mmu.Write(
		1,
		0,
		10,
	)

	if err != nil {
		panic(err)
	}

	err = mmu.Fork(
		1,
		2,
	)

	if err != nil {
		panic(err)
	}

	err = mmu.Write(
		2,
		0,
		99,
	)

	if err != nil {
		panic(err)
	}

	parent, _ := mmu.Read(
		1,
		0,
	)

	child, _ := mmu.Read(
		2,
		0,
	)

	fmt.Println(
		"parent:",
		parent,
	)

	fmt.Println(
		"child:",
		child,
	)

	fmt.Println(
		"refcount-parent:",
		mmu.RefCount(1),
	)

	fmt.Println(
		"refcount-child:",
	)
}

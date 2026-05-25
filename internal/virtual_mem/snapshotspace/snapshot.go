package snapshotspace

import (
	"branchkv-core/internal/virtual_mem/pagetable/l1"
)

type Space struct {
	ID     uint64
	Root   *l1.Table
	Parent *Space
}

func NewSpace(
	id uint64,
	root *l1.Table,
) *Space {

	return &Space{
		ID:   id,
		Root: root,
	}
}

func Fork(
	id uint64,
	parent *Space,
) *Space {

	return &Space{
		ID:     id,
		Root:   parent.Root,
		Parent: parent,
	}
}

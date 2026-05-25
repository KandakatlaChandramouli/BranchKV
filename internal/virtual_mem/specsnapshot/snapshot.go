package specsnapshot

import (
	"branchkv-core/internal/virtual_mem/pagetable/l1"
)

type Snapshot struct {
	ID   uint64
	Root *l1.Table
}

func NewSnapshot(
	id uint64,
	root *l1.Table,
) *Snapshot {

	return &Snapshot{
		ID:   id,
		Root: root,
	}
}

package concurrent

import (
	"branchkv-core/internal/tree"
	"sync"
)

type BranchTable struct {
	table sync.Map
}

func NewBranchTable() *BranchTable {
	return &BranchTable{}
}

func (b *BranchTable) Store(
	id uint64,
	branch *tree.Branch,
) {

	b.table.Store(id, branch)
}

func (b *BranchTable) Load(
	id uint64,
) (*tree.Branch, bool) {

	v, ok := b.table.Load(id)

	if !ok {
		return nil, false
	}

	return v.(*tree.Branch), true
}

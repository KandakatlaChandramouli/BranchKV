package specarena

import (
	"branchkv-core/internal/virtual_mem/fixeddag"
)

type Arena struct {
	nodes []fixeddag.Node
	next  uint64
	limit uint64
}

func NewArena(
	size uint64,
) *Arena {

	return &Arena{
		nodes: make(
			[]fixeddag.Node,
			size,
		),
		limit: size,
	}
}

func (a *Arena) Allocate() *fixeddag.Node {

	if a.next >= a.limit {
		return nil
	}

	index := a.next

	a.next++

	return &a.nodes[index]
}

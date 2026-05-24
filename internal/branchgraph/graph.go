package branchgraph

import (
	"sync"
)

type Edge struct {
	From uint64
	To   uint64
}

type BranchGraph struct {
	edges []Edge
	mu    sync.RWMutex
}

func NewBranchGraph() *BranchGraph {

	return &BranchGraph{
		edges: make([]Edge, 0),
	}
}

func (b *BranchGraph) Connect(
	from uint64,
	to uint64,
) {

	b.mu.Lock()
	defer b.mu.Unlock()

	b.edges = append(
		b.edges,
		Edge{
			From: from,
			To:   to,
		},
	)
}

func (b *BranchGraph) Size() int {

	b.mu.RLock()
	defer b.mu.RUnlock()

	return len(b.edges)
}

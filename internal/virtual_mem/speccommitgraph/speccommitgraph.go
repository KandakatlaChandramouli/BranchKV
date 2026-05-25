package speccommitgraph

import "sync"

type Graph struct {
	mu    sync.RWMutex
	edges map[uint64][]uint64
}

func NewGraph() *Graph {

	return &Graph{
		edges: make(
			map[uint64][]uint64,
		),
	}
}

func (g *Graph) Link(
	parent uint64,
	child uint64,
) {

	g.mu.Lock()
	defer g.mu.Unlock()

	g.edges[parent] = append(
		g.edges[parent],
		child,
	)
}

func (g *Graph) Children(
	parent uint64,
) []uint64 {

	g.mu.RLock()
	defer g.mu.RUnlock()

	return g.edges[parent]
}

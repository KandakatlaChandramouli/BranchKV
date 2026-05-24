package executiongraph

import (
	"sync"
)

type Node struct {
	ID uint64
}

type Edge struct {
	From uint64
	To   uint64
}

type ExecutionGraph struct {
	nodes []Node
	edges []Edge
	mu    sync.RWMutex
}

func NewExecutionGraph() *ExecutionGraph {

	return &ExecutionGraph{
		nodes: make([]Node, 0),
		edges: make([]Edge, 0),
	}
}

func (e *ExecutionGraph) AddNode(
	id uint64,
) {

	e.mu.Lock()
	defer e.mu.Unlock()

	e.nodes = append(
		e.nodes,
		Node{
			ID: id,
		},
	)
}

func (e *ExecutionGraph) AddEdge(
	from uint64,
	to uint64,
) {

	e.mu.Lock()
	defer e.mu.Unlock()

	e.edges = append(
		e.edges,
		Edge{
			From: from,
			To:   to,
		},
	)
}

func (e *ExecutionGraph) Nodes() int {

	e.mu.RLock()
	defer e.mu.RUnlock()

	return len(e.nodes)
}

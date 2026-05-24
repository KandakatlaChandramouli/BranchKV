package controlplane

import (
	"sync"
)

type RuntimeNode struct {
	ID     uint64
	Status string
}

type ControlPlane struct {
	nodes map[uint64]RuntimeNode
	mu    sync.RWMutex
}

func NewControlPlane() *ControlPlane {

	return &ControlPlane{
		nodes: make(
			map[uint64]RuntimeNode,
		),
	}
}

func (c *ControlPlane) Register(
	node RuntimeNode,
) {

	c.mu.Lock()
	defer c.mu.Unlock()

	c.nodes[node.ID] = node
}

func (c *ControlPlane) Size() int {

	c.mu.RLock()
	defer c.mu.RUnlock()

	return len(c.nodes)
}

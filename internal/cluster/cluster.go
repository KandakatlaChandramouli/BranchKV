package cluster

import (
	"sync"
)

type Node struct {
	ID      uint64
	Address string
}

type Cluster struct {
	nodes map[uint64]Node
	mu    sync.RWMutex
}

func NewCluster() *Cluster {

	return &Cluster{
		nodes: make(
			map[uint64]Node,
		),
	}
}

func (c *Cluster) AddNode(
	n Node,
) {

	c.mu.Lock()
	defer c.mu.Unlock()

	c.nodes[n.ID] = n
}

func (c *Cluster) Size() int {

	c.mu.RLock()
	defer c.mu.RUnlock()

	return len(c.nodes)
}

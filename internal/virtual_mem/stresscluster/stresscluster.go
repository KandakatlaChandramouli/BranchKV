package stresscluster

import "sync"

type Node struct {
	ID uint64
}

type Cluster struct {
	mu    sync.RWMutex
	nodes []Node
}

func NewCluster() *Cluster {

	return &Cluster{}
}

func (c *Cluster) Add(
	id uint64,
) {

	c.mu.Lock()
	defer c.mu.Unlock()

	c.nodes = append(
		c.nodes,
		Node{
			ID: id,
		},
	)
}

func (c *Cluster) Size() int {

	c.mu.RLock()
	defer c.mu.RUnlock()

	return len(
		c.nodes,
	)
}

package fabric

import (
	"sync"
)

type FabricNode struct {
	ID uint64
}

type Fabric struct {
	nodes []FabricNode
	mu    sync.RWMutex
}

func NewFabric() *Fabric {

	return &Fabric{
		nodes: make([]FabricNode, 0),
	}
}

func (f *Fabric) AddNode(
	id uint64,
) {

	f.mu.Lock()
	defer f.mu.Unlock()

	f.nodes = append(
		f.nodes,
		FabricNode{
			ID: id,
		},
	)
}

func (f *Fabric) Size() int {

	f.mu.RLock()
	defer f.mu.RUnlock()

	return len(f.nodes)
}

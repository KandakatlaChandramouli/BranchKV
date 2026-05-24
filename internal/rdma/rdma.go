package rdma

import (
	"sync/atomic"
)

type RDMARegion struct {
	ID   uint64
	Size uint64
}

type RDMAEngine struct {
	counter atomic.Uint64
}

func NewRDMAEngine() *RDMAEngine {
	return &RDMAEngine{}
}

func (r *RDMAEngine) Allocate(
	size uint64,
) *RDMARegion {

	id := r.counter.Add(1)

	return &RDMARegion{
		ID:   id,
		Size: size,
	}
}

package dma

import (
	"sync/atomic"
)

type DMARegion struct {
	ID   uint64
	Size uint64
}

type DMAEngine struct {
	counter atomic.Uint64
}

func NewDMAEngine() *DMAEngine {
	return &DMAEngine{}
}

func (d *DMAEngine) Allocate(
	size uint64,
) *DMARegion {

	id := d.counter.Add(1)

	return &DMARegion{
		ID:   id,
		Size: size,
	}
}

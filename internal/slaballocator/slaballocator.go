package slaballocator

import (
	"sync"
)

type Slab struct {
	objects [][]byte
}

type Allocator struct {
	slabs []Slab
	mu    sync.Mutex
}

func NewAllocator() *Allocator {

	return &Allocator{
		slabs: make([]Slab, 0),
	}
}

func (a *Allocator) Allocate(
	size int,
) []byte {

	a.mu.Lock()
	defer a.mu.Unlock()

	obj := make(
		[]byte,
		size,
	)

	a.slabs = append(
		a.slabs,
		Slab{
			objects: [][]byte{
				obj,
			},
		},
	)

	return obj
}

func (a *Allocator) Size() int {

	a.mu.Lock()
	defer a.mu.Unlock()

	return len(a.slabs)
}

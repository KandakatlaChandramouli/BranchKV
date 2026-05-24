package allocator_rt

import (
	"sync/atomic"
)

type RuntimeAllocator struct {
	allocs atomic.Uint64
}

func NewRuntimeAllocator() *RuntimeAllocator {

	return &RuntimeAllocator{}
}

func (r *RuntimeAllocator) Allocate() {
	r.allocs.Add(1)
}

func (r *RuntimeAllocator) Count() uint64 {
	return r.allocs.Load()
}

package tokenallocator

import (
	"sync/atomic"
)

type Allocator struct {
	tokens atomic.Uint64
}

func NewAllocator() *Allocator {

	return &Allocator{}
}

func (a *Allocator) Allocate() uint64 {
	return a.tokens.Add(1)
}

func (a *Allocator) Count() uint64 {
	return a.tokens.Load()
}

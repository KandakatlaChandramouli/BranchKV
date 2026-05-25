package eviction

import (
	"branchkv-core/internal/virtual_mem"
	"branchkv-core/internal/virtual_mem/swap"
)

type Evictor struct {
	swap *swap.Manager
}

func NewEvictor(
	swapper *swap.Manager,
) *Evictor {

	return &Evictor{
		swap: swapper,
	}
}

func (e *Evictor) Evict(
	frame *virtual_mem.PhysicalFrame,
) {

	if frame.Dirty.Load() {

		e.swap.SwapOut(
			frame,
		)
	}

	frame.InUse.Store(false)
}

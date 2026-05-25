package reclaimer

import (
	"branchkv-core/internal/virtual_mem"
)

type Reclaimer struct{}

func NewReclaimer() *Reclaimer {

	return &Reclaimer{}
}

func (r *Reclaimer) Reclaim(
	frame *virtual_mem.PhysicalFrame,
) bool {

	if frame.RefCount.Load() != 0 {
		return false
	}

	frame.InUse.Store(false)

	return true
}

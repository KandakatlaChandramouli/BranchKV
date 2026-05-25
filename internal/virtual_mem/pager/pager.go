package pager

import (
	"branchkv-core/internal/virtual_mem"
	"branchkv-core/internal/virtual_mem/swap"
)

type Pager struct {
	swap *swap.Manager
}

func NewPager(
	swapper *swap.Manager,
) *Pager {

	return &Pager{
		swap: swapper,
	}
}

func (p *Pager) Fault(
	frame *virtual_mem.PhysicalFrame,
) bool {

	return p.swap.SwapIn(
		frame,
	)
}

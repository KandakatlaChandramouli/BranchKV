package hypervisor

import (
	"sync/atomic"
)

type VM struct {
	ID uint64
}

type Hypervisor struct {
	counter atomic.Uint64
}

func NewHypervisor() *Hypervisor {
	return &Hypervisor{}
}

func (h *Hypervisor) CreateVM() *VM {

	id := h.counter.Add(1)

	return &VM{
		ID: id,
	}
}

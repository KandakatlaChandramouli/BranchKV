package hazard

import (
	"branchkv-core/internal/virtual_mem"
	"sync"
)

type Pointer struct {
	mu sync.RWMutex

	frame *virtual_mem.PhysicalFrame
}

func NewPointer() *Pointer {

	return &Pointer{}
}

func (p *Pointer) Protect(
	frame *virtual_mem.PhysicalFrame,
) {

	p.mu.Lock()
	defer p.mu.Unlock()

	p.frame = frame
}

func (p *Pointer) Load() *virtual_mem.PhysicalFrame {

	p.mu.RLock()
	defer p.mu.RUnlock()

	return p.frame
}

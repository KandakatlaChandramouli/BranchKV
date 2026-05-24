package virtual_mem

import (
	"sync/atomic"
)

const CacheLinePad = 64

type PhysicalPage struct {
	ID       uint64
	RefCount atomic.Int64
	Data     []float32

	_ [CacheLinePad]byte
}

func NewPhysicalPage(id uint64, size int) *PhysicalPage {

	p := &PhysicalPage{
		ID:   id,
		Data: make([]float32, size),
	}

	p.RefCount.Store(1)

	return p
}

func (p *PhysicalPage) IncRef() {
	p.RefCount.Add(1)
}

func (p *PhysicalPage) DecRef() int64 {
	return p.RefCount.Add(-1)
}

func (p *PhysicalPage) Count() int64 {
	return p.RefCount.Load()
}

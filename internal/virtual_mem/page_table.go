package virtual_mem

import "sync"

type PageTable struct {
	mu      sync.RWMutex
	entries map[uint64]*PhysicalFrame
}

func NewPageTable() *PageTable {

	return &PageTable{
		entries: make(
			map[uint64]*PhysicalFrame,
		),
	}
}

func (p *PageTable) Map(
	virtual uint64,
	frame *PhysicalFrame,
) {

	p.mu.Lock()
	defer p.mu.Unlock()

	p.entries[virtual] = frame
}

func (p *PageTable) Resolve(
	virtual uint64,
) (
	*PhysicalFrame,
	bool,
) {

	p.mu.RLock()
	defer p.mu.RUnlock()

	frame, ok := p.entries[virtual]

	return frame, ok
}

func (p *PageTable) Unmap(
	virtual uint64,
) {

	p.mu.Lock()
	defer p.mu.Unlock()

	delete(
		p.entries,
		virtual,
	)
}

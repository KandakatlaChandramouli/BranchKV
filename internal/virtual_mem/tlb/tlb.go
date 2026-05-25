package tlb

import (
	"branchkv-core/internal/virtual_mem"
	"sync"
)

type Cache struct {
	mu      sync.RWMutex
	entries map[uint64]*virtual_mem.PhysicalFrame
}

func NewCache() *Cache {

	return &Cache{
		entries: make(
			map[uint64]*virtual_mem.PhysicalFrame,
		),
	}
}

func (t *Cache) Lookup(
	virtual uint64,
) (
	*virtual_mem.PhysicalFrame,
	bool,
) {

	t.mu.RLock()
	defer t.mu.RUnlock()

	frame, ok := t.entries[virtual]

	return frame, ok
}

func (t *Cache) Insert(
	virtual uint64,
	frame *virtual_mem.PhysicalFrame,
) {

	t.mu.Lock()
	defer t.mu.Unlock()

	t.entries[virtual] = frame
}

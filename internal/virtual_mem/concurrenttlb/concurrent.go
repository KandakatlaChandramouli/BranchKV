package concurrenttlb

import (
	"branchkv-core/internal/virtual_mem"
	"sync"
)

type Cache struct {
	mu sync.RWMutex

	entries map[uint64]*virtual_mem.PhysicalFrame
}

func NewCache() *Cache {

	return &Cache{
		entries: make(
			map[uint64]*virtual_mem.PhysicalFrame,
		),
	}
}

func (c *Cache) Lookup(
	virtual uint64,
) (
	*virtual_mem.PhysicalFrame,
	bool,
) {

	c.mu.RLock()
	defer c.mu.RUnlock()

	frame, ok := c.entries[virtual]

	return frame, ok
}

func (c *Cache) Insert(
	virtual uint64,
	frame *virtual_mem.PhysicalFrame,
) {

	c.mu.Lock()
	defer c.mu.Unlock()

	c.entries[virtual] = frame
}

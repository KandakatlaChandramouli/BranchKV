package tablecache

import (
	"sync"
)

type Cache struct {
	tables map[string]bool
	mu     sync.RWMutex
}

func NewCache() *Cache {

	return &Cache{
		tables: make(
			map[string]bool,
		),
	}
}

func (c *Cache) Add(
	name string,
) {

	c.mu.Lock()
	defer c.mu.Unlock()

	c.tables[name] = true
}

func (c *Cache) Size() int {

	c.mu.RLock()
	defer c.mu.RUnlock()

	return len(c.tables)
}

package attentioncache

import (
	"sync"
)

type Cache struct {
	data map[uint64][]float32
	mu   sync.RWMutex
}

func NewCache() *Cache {

	return &Cache{
		data: make(
			map[uint64][]float32,
		),
	}
}

func (c *Cache) Store(
	id uint64,
	vec []float32,
) {

	c.mu.Lock()
	defer c.mu.Unlock()

	c.data[id] = vec
}

func (c *Cache) Size() int {

	c.mu.RLock()
	defer c.mu.RUnlock()

	return len(c.data)
}

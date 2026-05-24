package prefixcache

import (
	"sync"
)

type Cache struct {
	data map[string][]float32
	mu   sync.RWMutex
}

func NewCache() *Cache {

	return &Cache{
		data: make(
			map[string][]float32,
		),
	}
}

func (c *Cache) Store(
	key string,
	vec []float32,
) {

	c.mu.Lock()
	defer c.mu.Unlock()

	c.data[key] = vec
}

func (c *Cache) Size() int {

	c.mu.RLock()
	defer c.mu.RUnlock()

	return len(c.data)
}

package kernelcache

import (
	"sync"
)

type Kernel struct {
	Name string
}

type Cache struct {
	kernels map[string]Kernel
	mu      sync.RWMutex
}

func NewCache() *Cache {

	return &Cache{
		kernels: make(
			map[string]Kernel,
		),
	}
}

func (c *Cache) Store(
	name string,
) {

	c.mu.Lock()
	defer c.mu.Unlock()

	c.kernels[name] = Kernel{
		Name: name,
	}
}

func (c *Cache) Size() int {

	c.mu.RLock()
	defer c.mu.RUnlock()

	return len(c.kernels)
}

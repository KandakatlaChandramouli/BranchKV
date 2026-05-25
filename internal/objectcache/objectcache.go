package objectcache

import "sync"

type Cache struct {
	objects map[string][]byte
	mu      sync.RWMutex
}

func NewCache() *Cache {

	return &Cache{
		objects: make(
			map[string][]byte,
		),
	}
}

func (c *Cache) Put(
	key string,
	value []byte,
) {

	c.mu.Lock()
	defer c.mu.Unlock()

	c.objects[key] = value
}

func (c *Cache) Get(
	key string,
) (
	[]byte,
	bool,
) {

	c.mu.RLock()
	defer c.mu.RUnlock()

	v, ok := c.objects[key]

	return v, ok
}

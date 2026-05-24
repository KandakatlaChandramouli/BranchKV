package pagecache

import "sync"

type Cache struct {
	pages map[uint64][]byte
	mu    sync.RWMutex
}

func NewCache() *Cache {

	return &Cache{
		pages: make(
			map[uint64][]byte,
		),
	}
}

func (c *Cache) Put(
	id uint64,
	data []byte,
) {

	c.mu.Lock()
	defer c.mu.Unlock()

	c.pages[id] = data
}

func (c *Cache) Get(
	id uint64,
) (
	[]byte,
	bool,
) {

	c.mu.RLock()
	defer c.mu.RUnlock()

	v, ok := c.pages[id]

	return v, ok
}

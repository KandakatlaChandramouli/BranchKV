package pagecache

import (
	"sync"
)

type CachedPage struct {
	ID   uint64
	Data []byte
}

type PageCache struct {
	pages map[uint64]*CachedPage
	mu    sync.RWMutex
}

func NewPageCache() *PageCache {

	return &PageCache{
		pages: make(
			map[uint64]*CachedPage,
		),
	}
}

func (p *PageCache) Store(
	page *CachedPage,
) {

	p.mu.Lock()
	defer p.mu.Unlock()

	p.pages[page.ID] = page
}

func (p *PageCache) Load(
	id uint64,
) (*CachedPage, bool) {

	p.mu.RLock()
	defer p.mu.RUnlock()

	v, ok := p.pages[id]

	return v, ok
}

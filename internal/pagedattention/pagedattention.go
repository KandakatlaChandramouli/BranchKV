package pagedattention

import (
	"sync"
)

type AttentionPage struct {
	ID uint64
}

type PagedAttention struct {
	pages []AttentionPage
	mu    sync.Mutex
}

func NewPagedAttention() *PagedAttention {

	return &PagedAttention{
		pages: make([]AttentionPage, 0),
	}
}

func (p *PagedAttention) Allocate(
	id uint64,
) {

	p.mu.Lock()
	defer p.mu.Unlock()

	p.pages = append(
		p.pages,
		AttentionPage{
			ID: id,
		},
	)
}

func (p *PagedAttention) Size() int {

	p.mu.Lock()
	defer p.mu.Unlock()

	return len(p.pages)
}

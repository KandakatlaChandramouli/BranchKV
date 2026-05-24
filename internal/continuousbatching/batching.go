package continuousbatching

import (
	"sync"
)

type Request struct {
	ID uint64
}

type Batcher struct {
	requests []Request
	mu       sync.Mutex
}

func NewBatcher() *Batcher {

	return &Batcher{
		requests: make([]Request, 0),
	}
}

func (b *Batcher) Add(
	id uint64,
) {

	b.mu.Lock()
	defer b.mu.Unlock()

	b.requests = append(
		b.requests,
		Request{
			ID: id,
		},
	)
}

func (b *Batcher) Size() int {

	b.mu.Lock()
	defer b.mu.Unlock()

	return len(b.requests)
}

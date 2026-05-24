package buddyallocator

import (
	"sync"
)

type Block struct {
	Size int
}

type Allocator struct {
	blocks []Block
	mu     sync.Mutex
}

func NewAllocator() *Allocator {

	return &Allocator{
		blocks: make([]Block, 0),
	}
}

func (a *Allocator) Allocate(
	size int,
) {

	a.mu.Lock()
	defer a.mu.Unlock()

	a.blocks = append(
		a.blocks,
		Block{
			Size: size,
		},
	)
}

func (a *Allocator) Size() int {

	a.mu.Lock()
	defer a.mu.Unlock()

	return len(a.blocks)
}

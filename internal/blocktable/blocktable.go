package blocktable

import (
	"sync"
)

type Block struct {
	ID uint64
}

type Table struct {
	blocks map[uint64]Block
	mu     sync.RWMutex
}

func NewTable() *Table {

	return &Table{
		blocks: make(
			map[uint64]Block,
		),
	}
}

func (t *Table) Insert(
	id uint64,
) {

	t.mu.Lock()
	defer t.mu.Unlock()

	t.blocks[id] = Block{
		ID: id,
	}
}

func (t *Table) Size() int {

	t.mu.RLock()
	defer t.mu.RUnlock()

	return len(t.blocks)
}

package l2

import (
	"branchkv-core/internal/virtual_mem"
	"sync"
)

type Table struct {
	mu      sync.RWMutex
	entries map[uint64]*virtual_mem.PhysicalFrame
}

func NewTable() *Table {

	return &Table{
		entries: make(
			map[uint64]*virtual_mem.PhysicalFrame,
		),
	}
}

func (t *Table) Map(
	index uint64,
	frame *virtual_mem.PhysicalFrame,
) {

	t.mu.Lock()
	defer t.mu.Unlock()

	t.entries[index] = frame
}

func (t *Table) Resolve(
	index uint64,
) (
	*virtual_mem.PhysicalFrame,
	bool,
) {

	t.mu.RLock()
	defer t.mu.RUnlock()

	frame, ok := t.entries[index]

	return frame, ok
}

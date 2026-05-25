package l1

import (
	"branchkv-core/internal/virtual_mem/pagetable/l2"
	"sync"
)

type Table struct {
	mu      sync.RWMutex
	entries map[uint64]*l2.Table
}

func NewTable() *Table {

	return &Table{
		entries: make(
			map[uint64]*l2.Table,
		),
	}
}

func (t *Table) Resolve(
	index uint64,
) (
	*l2.Table,
	bool,
) {

	t.mu.RLock()
	defer t.mu.RUnlock()

	entry, ok := t.entries[index]

	return entry, ok
}

func (t *Table) Ensure(
	index uint64,
) *l2.Table {

	t.mu.Lock()
	defer t.mu.Unlock()

	entry, ok := t.entries[index]

	if ok {
		return entry
	}

	entry = l2.NewTable()

	t.entries[index] = entry

	return entry
}

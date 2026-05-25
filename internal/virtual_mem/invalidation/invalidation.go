package invalidation

import "sync"

type Table struct {
	mu      sync.RWMutex
	entries map[uint64]struct{}
}

func NewTable() *Table {

	return &Table{
		entries: make(
			map[uint64]struct{},
		),
	}
}

func (t *Table) Invalidate(
	virtual uint64,
) {

	t.mu.Lock()
	defer t.mu.Unlock()

	t.entries[virtual] = struct{}{}
}

func (t *Table) Invalid(
	virtual uint64,
) bool {

	t.mu.RLock()
	defer t.mu.RUnlock()

	_, ok := t.entries[virtual]

	return ok
}

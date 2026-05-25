package specpagetable

import "sync"

type Entry struct {
	Virtual  uint64
	Physical uint64
}

type Table struct {
	mu      sync.RWMutex
	entries map[uint64]uint64
}

func NewTable() *Table {

	return &Table{
		entries: make(
			map[uint64]uint64,
		),
	}
}

func (t *Table) Map(
	virtual uint64,
	physical uint64,
) {

	t.mu.Lock()
	defer t.mu.Unlock()

	t.entries[virtual] = physical
}

func (t *Table) Resolve(
	virtual uint64,
) (
	uint64,
	bool,
) {

	t.mu.RLock()
	defer t.mu.RUnlock()

	physical, ok := t.entries[virtual]

	return physical, ok
}

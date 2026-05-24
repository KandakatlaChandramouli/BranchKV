package syscall

import (
	"sync"
)

type Syscall func() error

type Table struct {
	entries map[uint64]Syscall
	mu      sync.RWMutex
}

func NewTable() *Table {

	return &Table{
		entries: make(
			map[uint64]Syscall,
		),
	}
}

func (t *Table) Register(
	id uint64,
	s Syscall,
) {

	t.mu.Lock()
	defer t.mu.Unlock()

	t.entries[id] = s
}

func (t *Table) Invoke(
	id uint64,
) error {

	t.mu.RLock()
	defer t.mu.RUnlock()

	fn := t.entries[id]

	if fn == nil {
		return nil
	}

	return fn()
}

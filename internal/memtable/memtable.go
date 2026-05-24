package memtable

import (
	"sync"
)

type Table struct {
	data map[string][]byte
	mu   sync.RWMutex
}

func NewTable() *Table {

	return &Table{
		data: make(
			map[string][]byte,
		),
	}
}

func (t *Table) Put(
	key string,
	value []byte,
) {

	t.mu.Lock()
	defer t.mu.Unlock()

	t.data[key] = value
}

func (t *Table) Size() int {

	t.mu.RLock()
	defer t.mu.RUnlock()

	return len(t.data)
}

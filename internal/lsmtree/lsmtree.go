package lsmtree

import (
	"sync"
)

type Entry struct {
	Key   string
	Value []byte
}

type LSMTree struct {
	data map[string][]byte
	mu   sync.RWMutex
}

func NewLSMTree() *LSMTree {

	return &LSMTree{
		data: make(
			map[string][]byte,
		),
	}
}

func (l *LSMTree) Put(
	key string,
	value []byte,
) {

	l.mu.Lock()
	defer l.mu.Unlock()

	l.data[key] = value
}

func (l *LSMTree) Get(
	key string,
) ([]byte, bool) {

	l.mu.RLock()
	defer l.mu.RUnlock()

	v, ok := l.data[key]

	return v, ok
}

func (l *LSMTree) Size() int {

	l.mu.RLock()
	defer l.mu.RUnlock()

	return len(l.data)
}

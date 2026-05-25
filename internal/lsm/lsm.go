package lsm

import "sync"

type Entry struct {
	Key   string
	Value []byte
}

type LSMTree struct {
	memtable map[string][]byte
	mu       sync.RWMutex
}

func NewLSMTree() *LSMTree {

	return &LSMTree{
		memtable: make(
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

	l.memtable[key] = value
}

func (l *LSMTree) Get(
	key string,
) (
	[]byte,
	bool,
) {

	l.mu.RLock()
	defer l.mu.RUnlock()

	v, ok := l.memtable[key]

	return v, ok
}

func (l *LSMTree) Size() int {

	l.mu.RLock()
	defer l.mu.RUnlock()

	return len(
		l.memtable,
	)
}

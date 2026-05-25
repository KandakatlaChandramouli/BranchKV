package storebuffer

import "sync"

type Entry struct {
	Address uint64
	Value   uint64
}

type Buffer struct {
	mu      sync.Mutex
	entries []Entry
}

func NewBuffer() *Buffer {

	return &Buffer{}
}

func (b *Buffer) Push(
	address uint64,
	value uint64,
) {

	b.mu.Lock()
	defer b.mu.Unlock()

	b.entries = append(
		b.entries,
		Entry{
			Address: address,
			Value:   value,
		},
	)
}

func (b *Buffer) Size() int {

	b.mu.Lock()
	defer b.mu.Unlock()

	return len(
		b.entries,
	)
}

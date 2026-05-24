package writebuffer

import (
	"sync"
)

type Buffer struct {
	data [][]byte
	mu   sync.Mutex
}

func NewBuffer() *Buffer {

	return &Buffer{
		data: make([][]byte, 0),
	}
}

func (b *Buffer) Append(
	v []byte,
) {

	b.mu.Lock()
	defer b.mu.Unlock()

	b.data = append(
		b.data,
		v,
	)
}

func (b *Buffer) Size() int {

	b.mu.Lock()
	defer b.mu.Unlock()

	return len(b.data)
}

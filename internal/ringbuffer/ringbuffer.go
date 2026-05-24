package ringbuffer

import (
	"sync/atomic"
)

type RingBuffer[T any] struct {
	buffer []T
	size   uint64

	head atomic.Uint64
	tail atomic.Uint64
}

func NewRingBuffer[T any](
	size uint64,
) *RingBuffer[T] {

	return &RingBuffer[T]{
		buffer: make([]T, size),
		size:   size,
	}
}

func (r *RingBuffer[T]) Push(
	v T,
) {

	tail := r.tail.Load()

	r.buffer[tail%r.size] = v

	r.tail.Add(1)
}

func (r *RingBuffer[T]) Pop() (T, bool) {

	var zero T

	head := r.head.Load()

	if head >= r.tail.Load() {
		return zero, false
	}

	v := r.buffer[head%r.size]

	r.head.Add(1)

	return v, true
}

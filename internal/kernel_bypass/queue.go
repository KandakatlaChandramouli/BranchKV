package kernel_bypass

import (
	"sync/atomic"
)

type Queue struct {
	head atomic.Uint64
	tail atomic.Uint64

	data [][]byte
	size uint64
}

func NewQueue(
	size uint64,
) *Queue {

	return &Queue{
		data: make(
			[][]byte,
			size,
		),
		size: size,
	}
}

func (q *Queue) Push(
	v []byte,
) {

	tail := q.tail.Load()

	q.data[tail%q.size] = v

	q.tail.Add(1)
}

func (q *Queue) Pop() ([]byte, bool) {

	head := q.head.Load()

	if head >= q.tail.Load() {
		return nil, false
	}

	v := q.data[head%q.size]

	q.head.Add(1)

	return v, true
}

package lockfreequeue

import (
	"sync/atomic"
)

type Queue struct {
	head atomic.Uint64
	tail atomic.Uint64
}

func NewQueue() *Queue {

	return &Queue{}
}

func (q *Queue) Push() {
	q.tail.Add(1)
}

func (q *Queue) Pop() bool {

	head := q.head.Load()
	tail := q.tail.Load()

	if head >= tail {
		return false
	}

	q.head.Add(1)

	return true
}

func (q *Queue) Size() uint64 {

	head := q.head.Load()
	tail := q.tail.Load()

	return tail - head
}

package runqueue

import (
	"sync"
)

type Queue struct {
	tasks []uint64
	mu    sync.Mutex
}

func NewQueue() *Queue {

	return &Queue{
		tasks: make([]uint64, 0),
	}
}

func (q *Queue) Push(
	id uint64,
) {

	q.mu.Lock()
	defer q.mu.Unlock()

	q.tasks = append(
		q.tasks,
		id,
	)
}

func (q *Queue) Pop() (uint64, bool) {

	q.mu.Lock()
	defer q.mu.Unlock()

	if len(q.tasks) == 0 {
		return 0, false
	}

	v := q.tasks[0]

	q.tasks = q.tasks[1:]

	return v, true
}

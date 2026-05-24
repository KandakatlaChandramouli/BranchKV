package taskqueue

import (
	"sync"
)

type Task struct {
	ID uint64
}

type Queue struct {
	tasks []Task
	mu    sync.Mutex
}

func NewQueue() *Queue {

	return &Queue{
		tasks: make([]Task, 0),
	}
}

func (q *Queue) Push(
	t Task,
) {

	q.mu.Lock()
	defer q.mu.Unlock()

	q.tasks = append(
		q.tasks,
		t,
	)
}

func (q *Queue) Size() int {

	q.mu.Lock()
	defer q.mu.Unlock()

	return len(q.tasks)
}

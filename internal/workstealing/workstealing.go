package workstealing

import (
	"sync"
)

type Task struct {
	ID uint64
}

type Scheduler struct {
	tasks []Task
	mu    sync.Mutex
}

func NewScheduler() *Scheduler {

	return &Scheduler{
		tasks: make([]Task, 0),
	}
}

func (s *Scheduler) Push(
	id uint64,
) {

	s.mu.Lock()
	defer s.mu.Unlock()

	s.tasks = append(
		s.tasks,
		Task{
			ID: id,
		},
	)
}

func (s *Scheduler) Steal() bool {

	s.mu.Lock()
	defer s.mu.Unlock()

	if len(s.tasks) == 0 {
		return false
	}

	s.tasks = s.tasks[1:]

	return true
}

func (s *Scheduler) Size() int {

	s.mu.Lock()
	defer s.mu.Unlock()

	return len(s.tasks)
}

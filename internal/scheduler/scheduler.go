package scheduler

import (
	"branchkv-core/internal/tree"
	"sync/atomic"
)

type Scheduler struct {
	queue atomic.Uint64
}

func NewScheduler() *Scheduler {
	return &Scheduler{}
}

func (s *Scheduler) Schedule(
	branch *tree.Branch,
) uint64 {

	return s.queue.Add(1)
}

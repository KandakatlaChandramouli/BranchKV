package gpu_scheduler

import (
	"sync/atomic"
)

type Scheduler struct {
	kernels atomic.Uint64
}

func NewScheduler() *Scheduler {

	return &Scheduler{}
}

func (s *Scheduler) Dispatch() {
	s.kernels.Add(1)
}

func (s *Scheduler) Count() uint64 {
	return s.kernels.Load()
}

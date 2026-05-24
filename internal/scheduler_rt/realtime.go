package scheduler_rt

import (
	"sync/atomic"
)

type RTScheduler struct {
	tasks atomic.Uint64
}

func NewRTScheduler() *RTScheduler {

	return &RTScheduler{}
}

func (r *RTScheduler) Schedule() {
	r.tasks.Add(1)
}

func (r *RTScheduler) Count() uint64 {
	return r.tasks.Load()
}

package streamexecutor

import (
	"sync/atomic"
)

type Executor struct {
	streams atomic.Uint64
}

func NewExecutor() *Executor {

	return &Executor{}
}

func (e *Executor) Dispatch() {
	e.streams.Add(1)
}

func (e *Executor) Count() uint64 {
	return e.streams.Load()
}

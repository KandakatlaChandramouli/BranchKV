package preemption

import (
	"sync/atomic"
)

type Engine struct {
	events atomic.Uint64
}

func NewEngine() *Engine {

	return &Engine{}
}

func (e *Engine) Preempt() {
	e.events.Add(1)
}

func (e *Engine) Count() uint64 {
	return e.events.Load()
}

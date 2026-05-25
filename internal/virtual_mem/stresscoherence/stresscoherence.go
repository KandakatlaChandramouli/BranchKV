package stresscoherence

import (
	"sync"
	"sync/atomic"
)

type Runtime struct {
	mu      sync.RWMutex
	counter atomic.Uint64
}

func NewRuntime() *Runtime {
	return &Runtime{}
}

func (r *Runtime) Advance() uint64 {

	return r.counter.Add(1)
}

func (r *Runtime) Load() uint64 {

	return r.counter.Load()
}

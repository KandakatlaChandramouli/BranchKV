package cowownership

import "sync/atomic"

type Runtime struct {
	counter atomic.Uint64
}

func NewRuntime() *Runtime {

	return &Runtime{}
}

func (r *Runtime) Next() uint64 {

	return r.counter.Add(1)
}

func (r *Runtime) Load() uint64 {

	return r.counter.Load()
}

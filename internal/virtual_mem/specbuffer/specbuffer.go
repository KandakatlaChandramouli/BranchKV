package specbuffer

import "sync/atomic"

type Runtime struct {
	value atomic.Uint64
}

func NewRuntime() *Runtime {

	return &Runtime{}
}

func (r *Runtime) Next() uint64 {

	return r.value.Add(1)
}

func (r *Runtime) Load() uint64 {

	return r.value.Load()
}

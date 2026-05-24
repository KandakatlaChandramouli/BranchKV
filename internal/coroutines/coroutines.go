package coroutines

import (
	"sync/atomic"
)

type Runtime struct {
	routines atomic.Uint64
}

func NewRuntime() *Runtime {

	return &Runtime{}
}

func (r *Runtime) Spawn() {
	r.routines.Add(1)
}

func (r *Runtime) Count() uint64 {
	return r.routines.Load()
}

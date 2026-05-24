package cuda

import (
	"sync/atomic"
)

type CUDADevice struct {
	ID uint64
}

type Runtime struct {
	kernels atomic.Uint64
}

func NewRuntime() *Runtime {

	return &Runtime{}
}

func (r *Runtime) Launch() {
	r.kernels.Add(1)
}

func (r *Runtime) Count() uint64 {
	return r.kernels.Load()
}

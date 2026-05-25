package runtimekernel_v2

type Runtime struct {
	count uint64
}

func NewRuntime() *Runtime {

	return &Runtime{}
}

func (r *Runtime) Add() {
	r.count++
}

func (r *Runtime) Size() uint64 {
	return r.count
}

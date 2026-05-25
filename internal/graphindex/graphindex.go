package graphindex

type Runtime struct {
	count int
}

func NewRuntime() *Runtime {

	return &Runtime{}
}

func (r *Runtime) Add() {
	r.count++
}

func (r *Runtime) Size() int {
	return r.count
}

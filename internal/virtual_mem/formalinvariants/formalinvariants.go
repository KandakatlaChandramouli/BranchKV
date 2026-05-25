package formalinvariants

import "sync"

type Invariant struct {
	Name string
}

type Runtime struct {
	mu         sync.RWMutex
	invariants []Invariant
}

func NewRuntime() *Runtime {

	return &Runtime{}
}

func (r *Runtime) Register(
	name string,
) {

	r.mu.Lock()
	defer r.mu.Unlock()

	r.invariants = append(
		r.invariants,
		Invariant{
			Name: name,
		},
	)
}

func (r *Runtime) Count() int {

	r.mu.RLock()
	defer r.mu.RUnlock()

	return len(
		r.invariants,
	)
}

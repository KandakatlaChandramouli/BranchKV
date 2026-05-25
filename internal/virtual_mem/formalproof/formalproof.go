package formalproof

import "sync"

type Proof struct {
	Name   string
	Passed bool
}

type Runtime struct {
	mu     sync.RWMutex
	proofs []Proof
}

func NewRuntime() *Runtime {

	return &Runtime{}
}

func (r *Runtime) Verify(
	name string,
) {

	r.mu.Lock()
	defer r.mu.Unlock()

	r.proofs = append(
		r.proofs,
		Proof{
			Name:   name,
			Passed: true,
		},
	)
}

func (r *Runtime) Count() int {

	r.mu.RLock()
	defer r.mu.RUnlock()

	return len(
		r.proofs,
	)
}

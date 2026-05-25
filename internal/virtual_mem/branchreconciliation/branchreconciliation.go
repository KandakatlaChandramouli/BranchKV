package branchreconciliation

import "sync"

type Delta struct {
	Branch uint64
	Parent uint64
}

type Runtime struct {
	mu     sync.RWMutex
	deltas []Delta
}

func NewRuntime() *Runtime {

	return &Runtime{}
}

func (r *Runtime) Merge(
	branch uint64,
	parent uint64,
) {

	r.mu.Lock()
	defer r.mu.Unlock()

	r.deltas = append(
		r.deltas,
		Delta{
			Branch: branch,
			Parent: parent,
		},
	)
}

func (r *Runtime) Size() int {

	r.mu.RLock()
	defer r.mu.RUnlock()

	return len(
		r.deltas,
	)
}

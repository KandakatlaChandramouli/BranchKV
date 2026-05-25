package distributedconsensus

import "sync"

type Vote struct {
	Node uint64
}

type Runtime struct {
	mu    sync.RWMutex
	votes []Vote
}

func NewRuntime() *Runtime {

	return &Runtime{}
}

func (r *Runtime) Commit(
	node uint64,
) {

	r.mu.Lock()
	defer r.mu.Unlock()

	r.votes = append(
		r.votes,
		Vote{
			Node: node,
		},
	)
}

func (r *Runtime) Count() int {

	r.mu.RLock()
	defer r.mu.RUnlock()

	return len(
		r.votes,
	)
}

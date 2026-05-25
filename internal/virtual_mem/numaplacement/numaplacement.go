package numaplacement

import "sync"

type Placement struct {
	Node uint64
	Page uint64
}

type Runtime struct {
	mu         sync.RWMutex
	placements map[uint64]Placement
}

func NewRuntime() *Runtime {

	return &Runtime{
		placements: make(
			map[uint64]Placement,
		),
	}
}

func (r *Runtime) Assign(
	page uint64,
	node uint64,
) {

	r.mu.Lock()
	defer r.mu.Unlock()

	r.placements[page] = Placement{
		Node: node,
		Page: page,
	}
}

func (r *Runtime) Resolve(
	page uint64,
) (
	Placement,
	bool,
) {

	r.mu.RLock()
	defer r.mu.RUnlock()

	value, ok := r.placements[page]

	return value, ok
}

package deterministicreplay

import "sync"

type Event struct {
	ID uint64
}

type Runtime struct {
	mu     sync.RWMutex
	events []Event
}

func NewRuntime() *Runtime {

	return &Runtime{}
}

func (r *Runtime) Append(
	id uint64,
) {

	r.mu.Lock()
	defer r.mu.Unlock()

	r.events = append(
		r.events,
		Event{
			ID: id,
		},
	)
}

func (r *Runtime) Count() int {

	r.mu.RLock()
	defer r.mu.RUnlock()

	return len(
		r.events,
	)
}

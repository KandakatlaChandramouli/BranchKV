package runtime_state

import (
	"sync/atomic"
)

type RuntimeState struct {
	Branches atomic.Uint64
	Tokens   atomic.Uint64
	Forks    atomic.Uint64
}

func NewRuntimeState() *RuntimeState {
	return &RuntimeState{}
}

func (r *RuntimeState) RecordBranch() {
	r.Branches.Add(1)
}

func (r *RuntimeState) RecordToken() {
	r.Tokens.Add(1)
}

func (r *RuntimeState) RecordFork() {
	r.Forks.Add(1)
}

package recovery

import (
	"sync/atomic"
)

type Recovery struct {
	events atomic.Uint64
}

func NewRecovery() *Recovery {

	return &Recovery{}
}

func (r *Recovery) Recover() {
	r.events.Add(1)
}

func (r *Recovery) Count() uint64 {
	return r.events.Load()
}

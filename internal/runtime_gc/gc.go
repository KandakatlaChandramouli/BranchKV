package runtime_gc

import (
	"sync/atomic"
)

type RuntimeGC struct {
	cycles atomic.Uint64
}

func NewRuntimeGC() *RuntimeGC {
	return &RuntimeGC{}
}

func (r *RuntimeGC) Collect() {
	r.cycles.Add(1)
}

func (r *RuntimeGC) Cycles() uint64 {
	return r.cycles.Load()
}

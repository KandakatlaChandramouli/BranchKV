package backpressure

import (
	"sync/atomic"
)

type Backpressure struct {
	dropped atomic.Uint64
}

func NewBackpressure() *Backpressure {

	return &Backpressure{}
}

func (b *Backpressure) Drop() {
	b.dropped.Add(1)
}

func (b *Backpressure) Count() uint64 {
	return b.dropped.Load()
}

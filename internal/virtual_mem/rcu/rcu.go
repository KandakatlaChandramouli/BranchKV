package rcu

import "sync/atomic"

type Epoch struct {
	value atomic.Uint64
}

func NewEpoch() *Epoch {

	return &Epoch{}
}

func (e *Epoch) Advance() uint64 {

	return e.value.Add(1)
}

func (e *Epoch) Load() uint64 {

	return e.value.Load()
}

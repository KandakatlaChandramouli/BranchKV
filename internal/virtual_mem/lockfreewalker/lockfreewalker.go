package lockfreewalker

import "sync/atomic"

type Walker struct {
	position atomic.Uint64
}

func NewWalker() *Walker {

	return &Walker{}
}

func (w *Walker) Advance() uint64 {

	return w.position.Add(1)
}

func (w *Walker) Position() uint64 {

	return w.position.Load()
}

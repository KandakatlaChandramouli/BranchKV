package epochgc

import (
	"sync/atomic"
)

type GC struct {
	epochs atomic.Uint64
}

func NewGC() *GC {

	return &GC{}
}

func (g *GC) Collect() {
	g.epochs.Add(1)
}

func (g *GC) Count() uint64 {
	return g.epochs.Load()
}

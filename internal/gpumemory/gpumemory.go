package gpumemory

import (
	"sync/atomic"
)

type Pool struct {
	allocated atomic.Uint64
}

func NewPool() *Pool {

	return &Pool{}
}

func (p *Pool) Allocate(
	bytes uint64,
) {

	p.allocated.Add(bytes)
}

func (p *Pool) Usage() uint64 {
	return p.allocated.Load()
}

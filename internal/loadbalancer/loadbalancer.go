package loadbalancer

import (
	"sync/atomic"
)

type Balancer struct {
	counter atomic.Uint64
}

func NewBalancer() *Balancer {

	return &Balancer{}
}

func (b *Balancer) Next(
	total uint64,
) uint64 {

	v := b.counter.Add(1)

	return v % total
}

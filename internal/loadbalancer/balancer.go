package loadbalancer

import (
	"sync/atomic"
)

type LoadBalancer struct {
	counter atomic.Uint64
}

func NewLoadBalancer() *LoadBalancer {
	return &LoadBalancer{}
}

func (l *LoadBalancer) Next(
	total int,
) int {

	v := l.counter.Add(1)

	return int(v % uint64(total))
}

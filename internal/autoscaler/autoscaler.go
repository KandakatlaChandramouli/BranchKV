package autoscaler

import (
	"sync/atomic"
)

type AutoScaler struct {
	nodes atomic.Uint64
}

func NewAutoScaler() *AutoScaler {

	return &AutoScaler{}
}

func (a *AutoScaler) ScaleUp() {
	a.nodes.Add(1)
}

func (a *AutoScaler) Nodes() uint64 {
	return a.nodes.Load()
}

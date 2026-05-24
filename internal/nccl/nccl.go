package nccl

import (
	"sync/atomic"
)

type NCCL struct {
	collectives atomic.Uint64
}

func NewNCCL() *NCCL {

	return &NCCL{}
}

func (n *NCCL) AllReduce() {
	n.collectives.Add(1)
}

func (n *NCCL) Count() uint64 {
	return n.collectives.Load()
}

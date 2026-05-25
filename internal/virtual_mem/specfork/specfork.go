package specfork

import (
	"sync/atomic"
)

type Branch struct {
	ID     uint64
	Parent uint64
}

type Forker struct {
	next atomic.Uint64
}

func NewForker() *Forker {

	return &Forker{}
}

func (f *Forker) Fork(
	parent uint64,
) Branch {

	return Branch{
		ID:     f.next.Add(1),
		Parent: parent,
	}
}

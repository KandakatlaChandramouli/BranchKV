package failover

import (
	"sync/atomic"
)

type Failover struct {
	events atomic.Uint64
}

func NewFailover() *Failover {

	return &Failover{}
}

func (f *Failover) Trigger() {
	f.events.Add(1)
}

func (f *Failover) Count() uint64 {
	return f.events.Load()
}

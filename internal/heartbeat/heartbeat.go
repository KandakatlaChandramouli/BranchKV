package heartbeat

import (
	"sync/atomic"
)

type Heartbeat struct {
	count atomic.Uint64
}

func NewHeartbeat() *Heartbeat {

	return &Heartbeat{}
}

func (h *Heartbeat) Beat() {
	h.count.Add(1)
}

func (h *Heartbeat) Count() uint64 {
	return h.count.Load()
}

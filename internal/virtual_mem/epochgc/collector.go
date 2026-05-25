package epochgc

import (
	"sync/atomic"
)

type Collector struct {
	epoch atomic.Uint64
}

func NewCollector() *Collector {

	return &Collector{}
}

func (c *Collector) Advance() uint64 {

	return c.epoch.Add(1)
}

func (c *Collector) Current() uint64 {

	return c.epoch.Load()
}

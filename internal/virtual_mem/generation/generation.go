package generation

import "sync/atomic"

type Counter struct {
	value atomic.Uint64
}

func NewCounter() *Counter {

	return &Counter{}
}

func (c *Counter) Next() uint64 {

	return c.value.Add(1)
}

func (c *Counter) Current() uint64 {

	return c.value.Load()
}

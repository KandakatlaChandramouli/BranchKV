package logshipping

import (
	"sync/atomic"
)

type Shipping struct {
	shipped atomic.Uint64
}

func NewShipping() *Shipping {

	return &Shipping{}
}

func (s *Shipping) Ship() {
	s.shipped.Add(1)
}

func (s *Shipping) Count() uint64 {
	return s.shipped.Load()
}

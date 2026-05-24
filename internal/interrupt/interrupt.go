package interrupt

import (
	"sync/atomic"
)

type InterruptController struct {
	count atomic.Uint64
}

func NewInterruptController() *InterruptController {

	return &InterruptController{}
}

func (i *InterruptController) Trigger() {
	i.count.Add(1)
}

func (i *InterruptController) Count() uint64 {
	return i.count.Load()
}

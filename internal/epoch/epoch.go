package epoch

import (
	"sync/atomic"
)

type EpochManager struct {
	current atomic.Uint64
}

func NewEpochManager() *EpochManager {
	return &EpochManager{}
}

func (e *EpochManager) Next() uint64 {
	return e.current.Add(1)
}

func (e *EpochManager) Current() uint64 {
	return e.current.Load()
}

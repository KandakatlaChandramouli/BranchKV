package process

import (
	"sync/atomic"
)

type Process struct {
	PID uint64
}

type ProcessManager struct {
	counter atomic.Uint64
}

func NewProcessManager() *ProcessManager {

	return &ProcessManager{}
}

func (p *ProcessManager) Spawn() *Process {

	id := p.counter.Add(1)

	return &Process{
		PID: id,
	}
}

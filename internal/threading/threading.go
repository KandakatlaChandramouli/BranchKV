package threading

import (
	"sync/atomic"
)

type Thread struct {
	TID uint64
}

type ThreadManager struct {
	counter atomic.Uint64
}

func NewThreadManager() *ThreadManager {

	return &ThreadManager{}
}

func (t *ThreadManager) Spawn() *Thread {

	id := t.counter.Add(1)

	return &Thread{
		TID: id,
	}
}

package kvpager

import (
	"sync"
)

type KVPage struct {
	ID   uint64
	Keys []float32
	Vals []float32
}

type KVPager struct {
	pages map[uint64]*KVPage
	mu    sync.RWMutex
}

func NewKVPager() *KVPager {

	return &KVPager{
		pages: make(
			map[uint64]*KVPage,
		),
	}
}

func (k *KVPager) Allocate(
	id uint64,
	size int,
) {

	k.mu.Lock()
	defer k.mu.Unlock()

	k.pages[id] = &KVPage{
		ID: id,
		Keys: make(
			[]float32,
			size,
		),
		Vals: make(
			[]float32,
			size,
		),
	}
}

func (k *KVPager) Size() int {

	k.mu.RLock()
	defer k.mu.RUnlock()

	return len(k.pages)
}

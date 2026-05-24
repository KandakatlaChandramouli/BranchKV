package kvstore_rt

import (
	"sync"
)

type KVStore struct {
	data map[uint64][]float32
	mu   sync.RWMutex
}

func NewKVStore() *KVStore {

	return &KVStore{
		data: make(
			map[uint64][]float32,
		),
	}
}

func (k *KVStore) Put(
	id uint64,
	vec []float32,
) {

	k.mu.Lock()
	defer k.mu.Unlock()

	k.data[id] = vec
}

func (k *KVStore) Get(
	id uint64,
) ([]float32, bool) {

	k.mu.RLock()
	defer k.mu.RUnlock()

	v, ok := k.data[id]

	return v, ok
}

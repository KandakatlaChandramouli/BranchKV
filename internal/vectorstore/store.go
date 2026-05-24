package vectorstore

import (
	"sync"
)

type VectorStore struct {
	data map[uint64][]float32
	mu   sync.RWMutex
}

func NewVectorStore() *VectorStore {

	return &VectorStore{
		data: make(map[uint64][]float32),
	}
}

func (v *VectorStore) Put(
	id uint64,
	vec []float32,
) {

	v.mu.Lock()
	defer v.mu.Unlock()

	v.data[id] = vec
}

func (v *VectorStore) Get(
	id uint64,
) ([]float32, bool) {

	v.mu.RLock()
	defer v.mu.RUnlock()

	val, ok := v.data[id]

	return val, ok
}

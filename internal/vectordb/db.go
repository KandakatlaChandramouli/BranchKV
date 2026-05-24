package vectordb

import (
	"math"
	"sync"
)

type VectorDB struct {
	data map[uint64][]float32
	mu   sync.RWMutex
}

func NewVectorDB() *VectorDB {

	return &VectorDB{
		data: make(
			map[uint64][]float32,
		),
	}
}

func (v *VectorDB) Insert(
	id uint64,
	vec []float32,
) {

	v.mu.Lock()
	defer v.mu.Unlock()

	v.data[id] = vec
}

func (v *VectorDB) Distance(
	a []float32,
	b []float32,
) float32 {

	var sum float64

	for i := 0; i < len(a); i++ {

		diff := float64(a[i] - b[i])

		sum += diff * diff
	}

	return float32(
		math.Sqrt(sum),
	)
}

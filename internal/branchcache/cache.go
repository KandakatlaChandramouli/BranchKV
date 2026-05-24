package branchcache

import (
	"branchkv-core/internal/types"
	"sync"
)

type BranchCache struct {
	data map[uint64][]types.TokenID
	mu   sync.RWMutex
}

func NewBranchCache() *BranchCache {

	return &BranchCache{
		data: make(
			map[uint64][]types.TokenID,
		),
	}
}

func (b *BranchCache) Store(
	id uint64,
	seq []types.TokenID,
) {

	b.mu.Lock()
	defer b.mu.Unlock()

	b.data[id] = seq
}

func (b *BranchCache) Load(
	id uint64,
) ([]types.TokenID, bool) {

	b.mu.RLock()
	defer b.mu.RUnlock()

	v, ok := b.data[id]

	return v, ok
}

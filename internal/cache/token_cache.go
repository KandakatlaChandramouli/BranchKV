package cache

import (
	"branchkv-core/internal/types"
	"sync"
)

type TokenCache struct {
	cache map[types.TokenID][]float32
	mu    sync.RWMutex
}

func NewTokenCache() *TokenCache {

	return &TokenCache{
		cache: make(map[types.TokenID][]float32),
	}
}

func (t *TokenCache) Store(
	token types.TokenID,
	vec []float32,
) {

	t.mu.Lock()
	defer t.mu.Unlock()

	t.cache[token] = vec
}

func (t *TokenCache) Load(
	token types.TokenID,
) ([]float32, bool) {

	t.mu.RLock()
	defer t.mu.RUnlock()

	v, ok := t.cache[token]

	return v, ok
}

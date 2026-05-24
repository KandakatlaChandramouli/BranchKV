package prefetch

import (
	"branchkv-core/internal/types"
)

type Prefetcher struct {
	window int
}

func NewPrefetcher(
	window int,
) *Prefetcher {

	return &Prefetcher{
		window: window,
	}
}

func (p *Prefetcher) Predict(
	tokens []types.TokenID,
) []types.TokenID {

	if len(tokens) <= p.window {
		return tokens
	}

	return tokens[len(tokens)-p.window:]
}

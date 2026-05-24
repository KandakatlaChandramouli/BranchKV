package pipeline

import (
	"branchkv-core/internal/types"
)

type Pipeline struct{}

func NewPipeline() *Pipeline {
	return &Pipeline{}
}

func (p *Pipeline) Process(
	seq []types.TokenID,
) []types.TokenID {

	out := make(
		[]types.TokenID,
		len(seq),
	)

	copy(out, seq)

	return out
}

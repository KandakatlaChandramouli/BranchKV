package speculative

import (
	"branchkv-core/internal/types"
)

type SpeculativeEngine struct{}

func NewSpeculativeEngine() *SpeculativeEngine {
	return &SpeculativeEngine{}
}

func (s *SpeculativeEngine) ForkSequence(
	seq []types.TokenID,
) [][]types.TokenID {

	out := make(
		[][]types.TokenID,
		0,
	)

	for i := 0; i < len(seq); i++ {

		cloned := make(
			[]types.TokenID,
			len(seq),
		)

		copy(cloned, seq)

		cloned = append(
			cloned,
			types.TokenID(i),
		)

		out = append(out, cloned)
	}

	return out
}

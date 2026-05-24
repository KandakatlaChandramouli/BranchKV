package specdecoder

import (
	"branchkv-core/internal/types"
)

type SpecDecoder struct{}

func NewSpecDecoder() *SpecDecoder {
	return &SpecDecoder{}
}

func (s *SpecDecoder) Predict(
	seq []types.TokenID,
) []types.TokenID {

	out := make(
		[]types.TokenID,
		len(seq),
	)

	copy(out, seq)

	out = append(
		out,
		types.TokenID(len(seq)+1),
	)

	return out
}

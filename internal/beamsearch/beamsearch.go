package beamsearch

import (
	"branchkv-core/internal/types"
)

type Beam struct {
	Tokens []types.TokenID
	Score  float32
}

type BeamSearch struct {
	Width int
}

func NewBeamSearch(
	width int,
) *BeamSearch {

	return &BeamSearch{
		Width: width,
	}
}

func (b *BeamSearch) Expand(
	seq []types.TokenID,
) []Beam {

	out := make(
		[]Beam,
		0,
	)

	for i := 0; i < b.Width; i++ {

		next := make(
			[]types.TokenID,
			len(seq),
		)

		copy(next, seq)

		next = append(
			next,
			types.TokenID(i),
		)

		out = append(
			out,
			Beam{
				Tokens: next,
				Score:  float32(i),
			},
		)
	}

	return out
}

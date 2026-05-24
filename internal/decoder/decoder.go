package decoder

import (
	"branchkv-core/internal/types"
)

type Decoder struct{}

func NewDecoder() *Decoder {
	return &Decoder{}
}

func (d *Decoder) Decode(
	seq []types.TokenID,
) types.TokenID {

	if len(seq) == 0 {
		return 0
	}

	return seq[len(seq)-1] + 1
}

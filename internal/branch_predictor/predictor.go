package branch_predictor

import (
	"branchkv-core/internal/types"
)

type Predictor struct{}

func NewPredictor() *Predictor {
	return &Predictor{}
}

func (p *Predictor) PredictNext(
	seq []types.TokenID,
) types.TokenID {

	if len(seq) == 0 {
		return 0
	}

	return seq[len(seq)-1]
}

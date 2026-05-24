package executor

import (
	"branchkv-core/internal/types"
)

type Executor struct{}

func NewExecutor() *Executor {
	return &Executor{}
}

func (e *Executor) Execute(
	seq []types.TokenID,
) types.TokenID {

	if len(seq) == 0 {
		return 0
	}

	return seq[len(seq)-1]
}

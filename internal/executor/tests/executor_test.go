package tests

import (
	"branchkv-core/internal/executor"
	"branchkv-core/internal/types"
	"testing"
)

func TestExecutor(
	t *testing.T,
) {

	e := executor.NewExecutor()

	v := e.Execute(
		[]types.TokenID{
			1, 2, 3,
		},
	)

	if v != 3 {
		t.Fatal("executor failed")
	}
}

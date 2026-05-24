package tests

import (
	"branchkv-core/internal/gpu"
	"testing"
)

func TestGPUExecute(
	t *testing.T,
) {

	engine := gpu.NewGPUExecutor(0)

	out := engine.Execute(
		[]float32{
			1, 2, 3,
		},
	)

	if len(out) != 3 {
		t.Fatal("gpu execute failed")
	}
}

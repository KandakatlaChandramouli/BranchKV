package tests

import (
	"branchkv-core/internal/cuda"
	"testing"
)

func TestCUDA(
	t *testing.T,
) {

	r := cuda.NewRuntime()

	r.Launch()

	if r.Count() != 1 {
		t.Fatal("cuda failed")
	}
}

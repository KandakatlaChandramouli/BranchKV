package tests

import (
	"branchkv-core/internal/operator_fusion"
	"testing"
)

func TestFusion(
	t *testing.T,
) {

	f := operator_fusion.NewFusionEngine()

	f.Fuse("matmul+relu")

	if f.Size() != 1 {
		t.Fatal("fusion failed")
	}
}

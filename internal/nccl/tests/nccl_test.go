package tests

import (
	"branchkv-core/internal/nccl"
	"testing"
)

func TestNCCL(
	t *testing.T,
) {

	n := nccl.NewNCCL()

	n.AllReduce()

	if n.Count() != 1 {
		t.Fatal("nccl failed")
	}
}

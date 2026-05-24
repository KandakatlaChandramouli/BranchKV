package tests

import (
	"branchkv-core/internal/rdma"
	"testing"
)

func TestRDMAAllocate(
	t *testing.T,
) {

	engine := rdma.NewRDMAEngine()

	region := engine.Allocate(4096)

	if region.Size != 4096 {
		t.Fatal("invalid rdma region")
	}
}

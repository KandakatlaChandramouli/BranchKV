package tests

import (
	"branchkv-core/internal/concurrent"
	"branchkv-core/internal/tree"
	"testing"
)

func TestBranchTable(
	t *testing.T,
) {

	bt := concurrent.NewBranchTable()

	b := &tree.Branch{
		ID: 1,
	}

	bt.Store(1, b)

	v, ok := bt.Load(1)

	if !ok {
		t.Fatal("load failed")
	}

	if v.ID != 1 {
		t.Fatal("invalid branch")
	}
}

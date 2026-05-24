package tests

import (
	"branchkv-core/internal/lsmtree"
	"testing"
)

func TestLSMTree(
	t *testing.T,
) {

	tree := lsmtree.NewLSMTree()

	tree.Put(
		"runtime",
		[]byte("branchkv"),
	)

	_, ok := tree.Get(
		"runtime",
	)

	if !ok {
		t.Fatal("lsm failed")
	}
}

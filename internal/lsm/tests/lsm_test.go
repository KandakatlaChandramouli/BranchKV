package tests

import (
	"branchkv-core/internal/lsm"
	"testing"
)

func TestLSM(
	t *testing.T,
) {

	tree := lsm.NewLSMTree()

	tree.Put(
		"branchkv",
		[]byte("runtime"),
	)

	_, ok := tree.Get(
		"branchkv",
	)

	if !ok {
		t.Fatal("lsm lookup failed")
	}
}

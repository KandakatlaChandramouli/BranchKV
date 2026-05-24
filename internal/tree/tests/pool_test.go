package tests

import (
	"branchkv-core/internal/tree"
	"testing"
)

func TestTreeCreation(t *testing.T) {

	rt := tree.NewRadixTree()

	if rt == nil {
		t.Fatal("tree creation failed")
	}
}

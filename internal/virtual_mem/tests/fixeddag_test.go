package tests

import (
	"branchkv-core/internal/virtual_mem/fixeddag"
	"branchkv-core/internal/virtual_mem/specarena"
	"testing"
)

func TestFixedDAG(
	t *testing.T,
) {

	arena := specarena.NewArena(
		8,
	)

	root := arena.Allocate()

	child := arena.Allocate()

	fixeddag.Link(
		root,
		child,
	)

	if root.Count != 1 {
		t.Fatal("invalid count")
	}
}

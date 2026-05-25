package tests

import (
	"branchkv-core/internal/virtual_mem/dagnodearena"
	"branchkv-core/internal/virtual_mem/specbranchgraph"
	"testing"
)

func TestArenaDAG(
	t *testing.T,
) {

	arena := dagnodearena.NewArena(
		16,
	)

	root := arena.Allocate()

	child := arena.Allocate()

	specbranchgraph.Link(
		root,
		child,
	)

	if child.Parent == nil {
		t.Fatal("missing parent")
	}
}

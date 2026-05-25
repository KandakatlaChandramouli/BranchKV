package tests

import (
	"branchkv-core/internal/virtual_mem/specdag"
	"testing"
)

func TestSpecDAG(
	t *testing.T,
) {

	root := specdag.NewNode(
		1,
	)

	child := specdag.NewNode(
		2,
	)

	specdag.Link(
		root,
		child,
	)

	if child.Parent == nil {
		t.Fatal("missing parent")
	}

	if len(root.Children) != 1 {
		t.Fatal("missing child")
	}
}

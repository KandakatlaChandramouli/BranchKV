package tests

import (
	"branchkv-core/internal/virtual_mem/pagetable/l1"
	"branchkv-core/internal/virtual_mem/rollback"
	"branchkv-core/internal/virtual_mem/snapshotspace"
	"testing"
)

func TestSnapshotRollback(
	t *testing.T,
) {

	root := l1.NewTable()

	parent := snapshotspace.NewSpace(
		1,
		root,
	)

	child := snapshotspace.Fork(
		2,
		parent,
	)

	if child.Parent == nil {
		t.Fatal("missing parent")
	}

	rollback.Restore(
		child,
		parent,
	)

	if child.Root != parent.Root {
		t.Fatal("rollback failed")
	}
}

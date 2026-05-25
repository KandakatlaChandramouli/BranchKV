package tests

import (
	"testing"

	"branchkv-core/internal/virtual_mem/branchreconciliation"
	"branchkv-core/internal/virtual_mem/deterministicreplay"
	"branchkv-core/internal/virtual_mem/formalinvariants"
	"branchkv-core/internal/virtual_mem/lockfreewalker"
	"branchkv-core/internal/virtual_mem/numaplacement"
)

func TestFinalRuntime(
	t *testing.T,
) {

	walker := lockfreewalker.NewWalker()

	if walker.Advance() == 0 {

		t.Fatal("walker failed")
	}

	numa := numaplacement.NewRuntime()

	numa.Assign(
		1,
		0,
	)

	if _, ok := numa.Resolve(
		1,
	); !ok {

		t.Fatal("numa resolve failed")
	}

	reconcile := branchreconciliation.NewRuntime()

	reconcile.Merge(
		10,
		1,
	)

	if reconcile.Size() == 0 {

		t.Fatal("reconcile failed")
	}

	replay := deterministicreplay.NewRuntime()

	replay.Append(
		1,
	)

	if replay.Count() == 0 {

		t.Fatal("replay failed")
	}

	invariants := formalinvariants.NewRuntime()

	invariants.Register(
		"cow-consistency",
	)

	if invariants.Count() == 0 {

		t.Fatal("invariant failed")
	}
}

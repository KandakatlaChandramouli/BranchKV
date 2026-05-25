package tests

import (
	"testing"

	"branchkv-core/internal/virtual_mem/distributedconsensus"
	"branchkv-core/internal/virtual_mem/formalproof"
	"branchkv-core/internal/virtual_mem/runtimeobservability"
	"branchkv-core/internal/virtual_mem/stresscluster"
)

func TestUltraFinalRuntime(
	t *testing.T,
) {

	proof := formalproof.NewRuntime()

	proof.Verify(
		"cow-invariant",
	)

	if proof.Count() == 0 {

		t.Fatal("formal proof failure")
	}

	cluster := stresscluster.NewCluster()

	cluster.Add(
		1,
	)

	if cluster.Size() == 0 {

		t.Fatal("cluster failure")
	}

	consensus := distributedconsensus.NewRuntime()

	consensus.Commit(
		1,
	)

	if consensus.Count() == 0 {

		t.Fatal("consensus failure")
	}

	observe := runtimeobservability.NewRuntime()

	observe.Record(
		1,
	)

	if observe.Count() == 0 {

		t.Fatal("observability failure")
	}
}

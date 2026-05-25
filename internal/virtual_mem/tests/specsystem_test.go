package tests

import (
	"testing"

	"branchkv-core/internal/virtual_mem/mmu"
	"branchkv-core/internal/virtual_mem/specfork"
	"branchkv-core/internal/virtual_mem/specmetrics"
	"branchkv-core/internal/virtual_mem/specpagetable"
)

func TestSpecSystem(
	t *testing.T,
) {

	table := specpagetable.NewTable()

	table.Map(
		0x1000,
		0x2000,
	)

	if _, ok := table.Resolve(
		0x1000,
	); !ok {

		t.Fatal("mapping failed")
	}

	memory := mmu.NewMMU()

	memory.Switch(
		&mmu.Context{
			ID: 1,
		},
	)

	if memory.Current() == nil {

		t.Fatal("mmu switch failed")
	}

	forker := specfork.NewForker()

	branch := forker.Fork(
		0,
	)

	if branch.ID == 0 {

		t.Fatal("fork failed")
	}

	metrics := specmetrics.NewMetrics()

	metrics.Commit()

	_, commits, _ := metrics.Snapshot()

	if commits == 0 {

		t.Fatal("metrics failed")
	}
}

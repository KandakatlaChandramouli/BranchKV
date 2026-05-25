package tests

import (
	"branchkv-core/internal/virtual_mem"
	"branchkv-core/internal/virtual_mem/epochgc"
	"branchkv-core/internal/virtual_mem/hazard"
	"testing"
)

func TestEpochGC(
	t *testing.T,
) {

	gc := epochgc.NewCollector()

	if gc.Advance() == 0 {
		t.Fatal("epoch failed")
	}
}

func TestHazardPointer(
	t *testing.T,
) {

	ptr := hazard.NewPointer()

	frame := virtual_mem.NewPhysicalFrame(
		1,
		64,
	)

	ptr.Protect(
		frame,
	)

	if ptr.Load() == nil {
		t.Fatal("hazard failed")
	}
}

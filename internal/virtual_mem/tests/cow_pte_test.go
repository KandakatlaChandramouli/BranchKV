package tests

import (
	"branchkv-core/internal/virtual_mem"
	"branchkv-core/internal/virtual_mem/cowfault"
	"branchkv-core/internal/virtual_mem/pte"
	"testing"
)

func TestPTEFlags(
	t *testing.T,
) {

	var entry pte.Entry

	entry.Set(
		pte.Present,
	)

	if !entry.Has(
		pte.Present,
	) {

		t.Fatal("missing flag")
	}
}

func TestCoWFault(
	t *testing.T,
) {

	frame := virtual_mem.NewPhysicalFrame(
		1,
		64,
	)

	frame.RefCount.Store(
		2,
	)

	next := cowfault.Handle(
		frame,
	)

	if next == nil {
		t.Fatal("clone failed")
	}
}

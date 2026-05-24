package tests

import (
	"branchkv-core/internal/slaballocator"
	"testing"
)

func TestSlabAllocator(
	t *testing.T,
) {

	a := slaballocator.NewAllocator()

	v := a.Allocate(128)

	if len(v) != 128 {
		t.Fatal("slab failed")
	}
}

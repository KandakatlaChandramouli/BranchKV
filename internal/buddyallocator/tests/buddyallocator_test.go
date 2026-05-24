package tests

import (
	"branchkv-core/internal/buddyallocator"
	"testing"
)

func TestBuddyAllocator(
	t *testing.T,
) {

	a := buddyallocator.NewAllocator()

	a.Allocate(256)

	if a.Size() != 1 {
		t.Fatal("buddy failed")
	}
}

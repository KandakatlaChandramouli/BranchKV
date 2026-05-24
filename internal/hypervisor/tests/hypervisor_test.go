package tests

import (
	"branchkv-core/internal/hypervisor"
	"testing"
)

func TestHypervisor(
	t *testing.T,
) {

	h := hypervisor.NewHypervisor()

	vm := h.CreateVM()

	if vm.ID == 0 {
		t.Fatal("vm failed")
	}
}

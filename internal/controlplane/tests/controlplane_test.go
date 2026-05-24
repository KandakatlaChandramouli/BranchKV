package tests

import (
	"branchkv-core/internal/controlplane"
	"testing"
)

func TestControlPlane(
	t *testing.T,
) {

	cp := controlplane.NewControlPlane()

	cp.Register(
		controlplane.RuntimeNode{
			ID:     1,
			Status: "healthy",
		},
	)

	if cp.Size() != 1 {
		t.Fatal("controlplane failed")
	}
}

package controlplane

import (
	"branchkv-core/internal/controlplane"
	"testing"
)

func BenchmarkControlPlane(
	b *testing.B,
) {

	cp := controlplane.NewControlPlane()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		cp.Register(
			controlplane.RuntimeNode{
				ID:     uint64(i),
				Status: "healthy",
			},
		)
	}
}

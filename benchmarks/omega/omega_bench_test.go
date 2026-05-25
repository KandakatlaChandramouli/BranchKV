package omega

import (
	"branchkv-core/internal/runtimefabric_v2"
	"testing"
)

func BenchmarkRuntimeFabricV2(
	b *testing.B,
) {

	rt := runtimefabric_v2.NewRuntime()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		rt.Add()
	}
}

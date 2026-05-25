package ultra

import (
	"branchkv-core/internal/runtimefabric"
	"testing"
)

func BenchmarkRuntimeFabric(
	b *testing.B,
) {

	rt := runtimefabric.NewRuntime()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		rt.Add()
	}
}

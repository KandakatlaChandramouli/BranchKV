package hyper

import (
	"branchkv-core/internal/runtimeorchestrator"
	"testing"
)

func BenchmarkRuntimeOrchestrator(
	b *testing.B,
) {

	rt := runtimeorchestrator.NewRuntime()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		rt.Add()
	}
}

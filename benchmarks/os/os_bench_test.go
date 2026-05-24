package os

import (
	"branchkv-core/internal/process"
	"branchkv-core/internal/threading"
	"testing"
)

func BenchmarkOSRuntime(
	b *testing.B,
) {

	pm := process.NewProcessManager()

	tm := threading.NewThreadManager()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		_ = pm.Spawn()

		_ = tm.Spawn()
	}
}

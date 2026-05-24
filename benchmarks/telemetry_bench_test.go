package benchmarks

import (
	"branchkv-core/internal/telemetry"
	"testing"
)

func BenchmarkTelemetry(
	b *testing.B,
) {

	tel := telemetry.NewTelemetry()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		tel.Record()
	}
}

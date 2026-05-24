package benchmarks

import (
	"branchkv-core/internal/snapshot"
	"testing"
)

type BenchState struct {
	Value int
	Data  []byte
}

func BenchmarkSnapshot(
	b *testing.B,
) {

	engine := snapshot.NewSnapshotEngine(
		"bench_snapshot.gob",
	)

	state := BenchState{
		Value: 42,
		Data:  make([]byte, 4096),
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = engine.Save(state)
	}
}

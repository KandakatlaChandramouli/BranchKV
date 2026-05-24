package benchmarks

import (
	"branchkv-core/internal/wal"
	"testing"
)

func BenchmarkWALAppend(
	b *testing.B,
) {

	w, err := wal.NewWAL(
		"bench.wal",
	)

	if err != nil {
		b.Fatal(err)
	}

	defer w.Close()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = w.Append(uint64(i))
	}
}

package storage

import (
	"branchkv-core/internal/wal"
	"testing"
)

func BenchmarkWALAppend(
	b *testing.B,
) {

	w, _ := wal.Open(
		"bench.wal",
	)

	defer w.Close()

	payload := []byte("branchkv")

	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		_ = w.Append(
			payload,
		)
	}
}

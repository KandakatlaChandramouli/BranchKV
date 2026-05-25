package lsm

import (
	"branchkv-core/internal/lsm"
	"testing"
)

func BenchmarkLSMPut(
	b *testing.B,
) {

	tree := lsm.NewLSMTree()

	payload := []byte("branchkv")

	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		tree.Put(
			"runtime",
			payload,
		)
	}
}

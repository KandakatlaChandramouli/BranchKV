package storage_engine

import (
	"branchkv-core/internal/lsmtree"
	"testing"
)

func BenchmarkStorageEngine(
	b *testing.B,
) {

	tree := lsmtree.NewLSMTree()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		tree.Put(
			"runtime",
			[]byte("value"),
		)
	}
}

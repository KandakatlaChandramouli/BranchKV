package storage

import (
	"branchkv-core/internal/objectstore"
	"branchkv-core/internal/sstable"
	"testing"
)

func BenchmarkStorageRuntime(
	b *testing.B,
) {

	store := objectstore.NewObjectStore()

	table := sstable.NewSSTable()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		store.Put(
			"key",
			[]byte("value"),
		)

		table.Append(
			"key",
			[]byte("value"),
		)
	}
}

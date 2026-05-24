package kernel

import (
	"branchkv-core/internal/runtime_gc"
	"branchkv-core/internal/vectordb"
	"testing"
)

func BenchmarkKernelRuntime(
	b *testing.B,
) {

	gc := runtime_gc.NewRuntimeGC()

	db := vectordb.NewVectorDB()

	db.Insert(
		1,
		[]float32{
			1, 2, 3, 4,
		},
	)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		gc.Collect()

		_ = db.Distance(
			[]float32{1, 2},
			[]float32{1, 2},
		)
	}
}

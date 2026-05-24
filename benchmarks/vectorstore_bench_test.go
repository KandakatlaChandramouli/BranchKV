package benchmarks

import (
	"branchkv-core/internal/vectorstore"
	"testing"
)

func BenchmarkVectorStore(
	b *testing.B,
) {

	vs := vectorstore.NewVectorStore()

	vec := []float32{
		1, 2, 3, 4,
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		vs.Put(
			uint64(i),
			vec,
		)

		_, _ = vs.Get(uint64(i))
	}
}

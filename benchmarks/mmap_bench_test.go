package benchmarks

import (
	"branchkv-core/internal/mmap"
	"testing"
)

func BenchmarkMMapWrite(
	b *testing.B,
) {

	region, err := mmap.NewMMapRegion(
		"bench.mmap",
		4096,
	)

	if err != nil {
		b.Fatal(err)
	}

	defer region.Close()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		region.Data[0] = byte(i)
	}
}

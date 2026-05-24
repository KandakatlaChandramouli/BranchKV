package benchmarks

import (
	"branchkv-core/internal/tree"
	"branchkv-core/internal/types"
	"branchkv-core/internal/virtual_mem"
	"testing"
)

func BenchmarkTrieInsert(
	b *testing.B,
) {

	rt := tree.NewRadixTree()

	mmu := virtual_mem.NewMMU()

	desc, _ := mmu.AllocateDescriptor(
		1,
		128,
	)

	seq := []types.TokenID{
		1, 2, 3, 4,
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		rt.Insert(
			seq,
			desc,
		)
	}
}

func BenchmarkTrieSearch(
	b *testing.B,
) {

	rt := tree.NewRadixTree()

	mmu := virtual_mem.NewMMU()

	desc, _ := mmu.AllocateDescriptor(
		1,
		128,
	)

	seq := []types.TokenID{
		1, 2, 3, 4,
	}

	rt.Insert(
		seq,
		desc,
	)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		_, _ = rt.Search(
			seq,
		)
	}
}

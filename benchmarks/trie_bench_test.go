package benchmarks

import (
	"branchkv-core/internal/tree"
	"branchkv-core/internal/types"
	"branchkv-core/internal/virtual_mem"
	"testing"
)

func BenchmarkTrieInsert(b *testing.B) {

	rt := tree.NewRadixTree()

	mmu := virtual_mem.NewMMU()

	seq := []types.TokenID{
		1, 2, 3, 4, 5, 6, 7, 8,
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		desc := mmu.Allocate(uint64(i+1), 512)

		rt.Insert(seq, desc)
	}
}

func BenchmarkTrieSearch(b *testing.B) {

	rt := tree.NewRadixTree()

	mmu := virtual_mem.NewMMU()

	seq := []types.TokenID{
		1, 2, 3, 4, 5, 6, 7, 8,
	}

	desc := mmu.Allocate(1, 512)

	rt.Insert(seq, desc)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		rt.Search(seq)
	}
}

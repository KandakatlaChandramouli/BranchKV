package benchmarks

import (
	"branchkv-core/internal/tree"
	"branchkv-core/internal/types"
	"testing"
)

func BenchmarkCompactTrieSearch(
	b *testing.B,
) {

	trie := tree.NewCompactTrie()

	seq := []types.TokenID{
		1, 2, 3, 4, 5, 6, 7, 8,
	}

	trie.Insert(seq)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		trie.Search(seq)
	}
}

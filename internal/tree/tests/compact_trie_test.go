package tests

import (
	"branchkv-core/internal/tree"
	"branchkv-core/internal/types"
	"testing"
)

func TestCompactTrie(t *testing.T) {

	trie := tree.NewCompactTrie()

	seq := []types.TokenID{
		1, 2, 3, 4,
	}

	trie.Insert(seq)

	found := trie.Search(seq)

	if !found {
		t.Fatal("compact trie search failed")
	}
}

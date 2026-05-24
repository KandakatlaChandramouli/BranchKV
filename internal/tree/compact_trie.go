package tree

import (
	"branchkv-core/internal/types"
)

type CompactTrie struct {
	Root *CompactNode
}

func NewCompactTrie() *CompactTrie {

	return &CompactTrie{
		Root: NewCompactNode(),
	}
}

func (t *CompactTrie) Insert(
	tokens []types.TokenID,
) {

	current := t.Root

	for _, tok := range tokens {

		current.mu.Lock()

		child := current.FindChild(tok)

		if child == nil {

			child = NewCompactNode()

			current.AddChild(tok, child)
		}

		current.mu.Unlock()

		current = child
	}

	current.End = true
}

func (t *CompactTrie) Search(
	tokens []types.TokenID,
) bool {

	current := t.Root

	for _, tok := range tokens {

		current.mu.RLock()

		child := current.FindChild(tok)

		current.mu.RUnlock()

		if child == nil {
			return false
		}

		current = child
	}

	return current.End
}

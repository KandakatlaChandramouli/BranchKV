package tree

import (
	"branchkv-core/internal/types"
	"branchkv-core/internal/virtual_mem"
)

type Branch struct {
	ID          uint64
	Tokens      []types.TokenID
	CurrentNode *Node
}

func NewBranch(
	id uint64,
	tokens []types.TokenID,
	node *Node,
) *Branch {

	return &Branch{
		ID:          id,
		Tokens:      tokens,
		CurrentNode: node,
	}
}

func ForkBranch(
	mmu *virtual_mem.MMU,
	parent *Branch,
	newID uint64,
) (*Branch, error) {

	err := mmu.Fork(parent.ID, newID)
	if err != nil {
		return nil, err
	}

	newTokens := make([]types.TokenID, len(parent.Tokens))

	copy(newTokens, parent.Tokens)

	return &Branch{
		ID:          newID,
		Tokens:      newTokens,
		CurrentNode: parent.CurrentNode,
	}, nil
}

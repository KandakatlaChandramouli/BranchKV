package tree

import (
	"branchkv-core/internal/types"
	"sync"
)

const MaxChildren = 16

type CompactChild struct {
	Token types.TokenID
	Node  *CompactNode
}

type CompactNode struct {
	Children [MaxChildren]CompactChild
	Count    int

	End bool

	mu sync.RWMutex
}

func NewCompactNode() *CompactNode {
	return &CompactNode{}
}

func (n *CompactNode) FindChild(
	token types.TokenID,
) *CompactNode {

	for i := 0; i < n.Count; i++ {

		if n.Children[i].Token == token {
			return n.Children[i].Node
		}
	}

	return nil
}

func (n *CompactNode) AddChild(
	token types.TokenID,
	child *CompactNode,
) {

	n.Children[n.Count] = CompactChild{
		Token: token,
		Node:  child,
	}

	n.Count++
}

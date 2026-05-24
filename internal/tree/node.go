package tree

import (
	"branchkv-core/internal/types"
	"branchkv-core/internal/virtual_mem"
	"sync"
)

type Node struct {
	Token      types.TokenID
	Children   map[types.TokenID]*Node
	Parent     *Node
	Descriptor *virtual_mem.VirtualDescriptor
	End        bool

	mu sync.RWMutex
}

func NewNode(token types.TokenID) *Node {
	return &Node{
		Token:    token,
		Children: make(map[types.TokenID]*Node),
	}
}

func (n *Node) Reset() {
	n.Token = 0
	n.Parent = nil
	n.Descriptor = nil
	n.End = false

	for k := range n.Children {
		delete(n.Children, k)
	}
}

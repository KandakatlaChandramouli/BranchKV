package tree

import (
	"branchkv-core/internal/virtual_mem"
)

type Node struct {
	ID         uint64
	Descriptor virtual_mem.VirtualDescriptor
	Children   []*Node
}

func NewNode(
	id uint64,
	desc virtual_mem.VirtualDescriptor,
) *Node {

	return &Node{
		ID:         id,
		Descriptor: desc,
		Children:   make([]*Node, 0),
	}
}

func (n *Node) AddChild(
	child *Node,
) {

	n.Children = append(
		n.Children,
		child,
	)
}

package tree

import (
	"branchkv-core/internal/pool"
	"branchkv-core/internal/types"
	"branchkv-core/internal/virtual_mem"
)

type RadixTree struct {
	Root     *Node
	nodePool *pool.GenericPool
}

func NewRadixTree() *RadixTree {

	p := pool.NewGenericPool(func() any {
		return NewNode(0)
	})

	return &RadixTree{
		Root:     NewNode(0),
		nodePool: p,
	}
}

func (r *RadixTree) acquireNode() *Node {
	n := r.nodePool.Get().(*Node)

	n.Reset()

	return n
}

func (r *RadixTree) releaseNode(n *Node) {
	r.nodePool.Put(n)
}

func (r *RadixTree) Insert(
	tokens []types.TokenID,
	desc *virtual_mem.VirtualDescriptor,
) *Node {

	current := r.Root

	for _, tok := range tokens {

		current.mu.Lock()

		child, exists := current.Children[tok]

		if !exists {

			child = r.acquireNode()

			child.Token = tok
			child.Parent = current

			current.Children[tok] = child
		}

		current.mu.Unlock()

		current = child
	}

	current.End = true
	current.Descriptor = desc

	return current
}

func (r *RadixTree) Search(
	tokens []types.TokenID,
) (*Node, bool) {

	current := r.Root

	for _, tok := range tokens {

		current.mu.RLock()

		child, exists := current.Children[tok]

		current.mu.RUnlock()

		if !exists {
			return nil, false
		}

		current = child
	}

	return current, current.End
}

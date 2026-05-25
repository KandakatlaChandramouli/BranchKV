package dagnodearena

type Node struct {
	ID       uint64
	Parent   *Node
	Children []*Node
}

type Arena struct {
	nodes []Node
	next  uint64
}

func NewArena(
	size uint64,
) *Arena {

	return &Arena{
		nodes: make(
			[]Node,
			size,
		),
	}
}

func (a *Arena) Allocate() *Node {

	index := a.next

	a.next++

	return &a.nodes[index]
}

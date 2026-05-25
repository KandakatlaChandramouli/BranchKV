package specdag

type Node struct {
	ID       uint64
	Parent   *Node
	Children []*Node
}

func NewNode(
	id uint64,
) *Node {

	return &Node{
		ID: id,
	}
}

func Link(
	parent *Node,
	child *Node,
) {

	child.Parent = parent

	parent.Children = append(
		parent.Children,
		child,
	)
}

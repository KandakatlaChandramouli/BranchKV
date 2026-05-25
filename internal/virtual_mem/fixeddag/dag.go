package fixeddag

const MaxChildren = 8

type Node struct {
	ID uint64

	Parent *Node

	Children [MaxChildren]*Node

	Count uint8
}

func Link(
	parent *Node,
	child *Node,
) {

	if parent.Count >= MaxChildren {
		return
	}

	index := parent.Count

	parent.Children[index] = child

	parent.Count++

	child.Parent = parent
}

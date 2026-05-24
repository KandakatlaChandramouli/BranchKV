package numa

type Node struct {
	ID uint64
}

type NUMA struct {
	nodes []Node
}

func NewNUMA() *NUMA {

	return &NUMA{
		nodes: make([]Node, 0),
	}
}

func (n *NUMA) AddNode(
	id uint64,
) {

	n.nodes = append(
		n.nodes,
		Node{
			ID: id,
		},
	)
}

func (n *NUMA) Size() int {
	return len(n.nodes)
}

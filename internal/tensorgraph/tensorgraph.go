package tensorgraph

type Node struct {
	ID uint64
}

type TensorGraph struct {
	nodes []Node
}

func NewTensorGraph() *TensorGraph {

	return &TensorGraph{
		nodes: make([]Node, 0),
	}
}

func (t *TensorGraph) AddNode(
	id uint64,
) {

	t.nodes = append(
		t.nodes,
		Node{
			ID: id,
		},
	)
}

func (t *TensorGraph) Size() int {
	return len(t.nodes)
}

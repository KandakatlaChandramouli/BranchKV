package dag

type Vertex struct {
	ID uint64
}

type DAG struct {
	vertices []Vertex
}

func NewDAG() *DAG {

	return &DAG{
		vertices: make([]Vertex, 0),
	}
}

func (d *DAG) AddVertex(
	id uint64,
) {

	d.vertices = append(
		d.vertices,
		Vertex{
			ID: id,
		},
	)
}

func (d *DAG) Size() int {
	return len(d.vertices)
}

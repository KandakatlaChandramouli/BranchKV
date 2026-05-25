package vectorindex

type Vector struct {
	ID     uint64
	Values []float32
}

type Index struct {
	vectors []Vector
}

func NewIndex() *Index {

	return &Index{
		vectors: make(
			[]Vector,
			0,
		),
	}
}

func (i *Index) Insert(
	id uint64,
	values []float32,
) {

	i.vectors = append(
		i.vectors,
		Vector{
			ID:     id,
			Values: values,
		},
	)
}

func (i *Index) Size() int {
	return len(i.vectors)
}

package versionvector

type Vector struct {
	values map[uint64]uint64
}

func NewVector() *Vector {

	return &Vector{
		values: make(
			map[uint64]uint64,
		),
	}
}

func (v *Vector) Update(
	node uint64,
	version uint64,
) {

	v.values[node] = version
}

func (v *Vector) Version(
	node uint64,
) uint64 {

	return v.values[node]
}

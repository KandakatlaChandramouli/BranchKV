package segmenttree

type Segment struct {
	ID uint64
}

type Tree struct {
	segments []Segment
}

func NewTree() *Tree {

	return &Tree{
		segments: make(
			[]Segment,
			0,
		),
	}
}

func (t *Tree) Add(
	id uint64,
) {

	t.segments = append(
		t.segments,
		Segment{
			ID: id,
		},
	)
}

func (t *Tree) Size() int {
	return len(t.segments)
}

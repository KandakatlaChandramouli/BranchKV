package compaction

type Segment struct {
	ID uint64
}

type Compactor struct{}

func NewCompactor() *Compactor {
	return &Compactor{}
}

func (c *Compactor) Compact(
	segments []Segment,
) Segment {

	if len(segments) == 0 {
		return Segment{}
	}

	return Segment{
		ID: segments[0].ID,
	}
}

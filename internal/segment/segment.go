package segment

type Segment struct {
	ID   uint64
	Data []byte
}

type SegmentManager struct {
	segments []Segment
}

func NewSegmentManager() *SegmentManager {

	return &SegmentManager{
		segments: make([]Segment, 0),
	}
}

func (s *SegmentManager) Append(
	seg Segment,
) {

	s.segments = append(
		s.segments,
		seg,
	)
}

func (s *SegmentManager) Size() int {
	return len(s.segments)
}

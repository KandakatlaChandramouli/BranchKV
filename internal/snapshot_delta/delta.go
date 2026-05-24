package snapshot_delta

type Delta struct {
	BaseVersion uint64
	Changes     []byte
}

type DeltaEngine struct{}

func NewDeltaEngine() *DeltaEngine {
	return &DeltaEngine{}
}

func (d *DeltaEngine) CreateDelta(
	base uint64,
	data []byte,
) *Delta {

	return &Delta{
		BaseVersion: base,
		Changes:     data,
	}
}

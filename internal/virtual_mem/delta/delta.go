package delta

type Delta struct {
	PageID uint64
	Offset uint64
	Old    float32
	New    float32
}

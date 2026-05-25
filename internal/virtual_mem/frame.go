package virtual_mem

import "sync/atomic"

type PhysicalFrame struct {
	ID       uint64
	RefCount atomic.Int64
	Dirty    atomic.Bool
	InUse    atomic.Bool
	Data     []float32
}

func NewPhysicalFrame(
	id uint64,
	size int,
) *PhysicalFrame {

	frame := &PhysicalFrame{
		ID: id,
		Data: make(
			[]float32,
			size,
		),
	}

	frame.RefCount.Store(1)
	frame.InUse.Store(true)

	return frame
}

func (f *PhysicalFrame) Retain() {
	f.RefCount.Add(1)
}

func (f *PhysicalFrame) Release() int64 {
	return f.RefCount.Add(-1)
}

func (f *PhysicalFrame) MarkDirty() {
	f.Dirty.Store(true)
}

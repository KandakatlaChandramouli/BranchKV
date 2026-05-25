package virtual_mem

type FrameAllocator struct {
	arena *Arena
}

func NewFrameAllocator(
	total uint64,
	frameSize int,
) *FrameAllocator {

	return &FrameAllocator{
		arena: NewArena(
			total,
			frameSize,
		),
	}
}

func (f *FrameAllocator) Alloc() (
	*PhysicalFrame,
	bool,
) {

	return f.arena.Allocate()
}

func (f *FrameAllocator) Free(
	frame *PhysicalFrame,
) {

	f.arena.Free(frame)
}

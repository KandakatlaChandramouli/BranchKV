package virtual_mem

import "sync"

type Arena struct {
	mu        sync.Mutex
	frameSize int
	frames    []*PhysicalFrame
	freeList  []uint64
}

func NewArena(
	total uint64,
	frameSize int,
) *Arena {

	arena := &Arena{
		frameSize: frameSize,
		frames: make(
			[]*PhysicalFrame,
			total,
		),
		freeList: make(
			[]uint64,
			0,
			total,
		),
	}

	for i := uint64(0); i < total; i++ {

		frame := NewPhysicalFrame(
			i,
			frameSize,
		)

		frame.InUse.Store(false)

		arena.frames[i] = frame

		arena.freeList = append(
			arena.freeList,
			i,
		)
	}

	return arena
}

func (a *Arena) Allocate() (
	*PhysicalFrame,
	bool,
) {

	a.mu.Lock()
	defer a.mu.Unlock()

	if len(a.freeList) == 0 {
		return nil, false
	}

	last := len(a.freeList) - 1

	id := a.freeList[last]

	a.freeList = a.freeList[:last]

	frame := a.frames[id]

	frame.InUse.Store(true)
	frame.RefCount.Store(1)

	return frame, true
}

func (a *Arena) Free(
	frame *PhysicalFrame,
) {

	a.mu.Lock()
	defer a.mu.Unlock()

	frame.InUse.Store(false)

	a.freeList = append(
		a.freeList,
		frame.ID,
	)
}

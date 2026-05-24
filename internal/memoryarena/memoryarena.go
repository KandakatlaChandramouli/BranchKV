package memoryarena

type Arena struct {
	buffer []byte
	offset int
}

func NewArena(
	size int,
) *Arena {

	return &Arena{
		buffer: make(
			[]byte,
			size,
		),
	}
}

func (a *Arena) Allocate(
	size int,
) []byte {

	start := a.offset

	end := start + size

	a.offset = end

	return a.buffer[start:end]
}

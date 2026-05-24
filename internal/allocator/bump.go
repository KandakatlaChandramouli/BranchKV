package allocator

type BumpAllocator struct {
	buffer []byte
	offset int
}

func NewBumpAllocator(
	size int,
) *BumpAllocator {

	return &BumpAllocator{
		buffer: make([]byte, size),
	}
}

func (b *BumpAllocator) Allocate(
	size int,
) []byte {

	start := b.offset

	end := start + size

	if end >= len(b.buffer) {
		return nil
	}

	b.offset = end

	return b.buffer[start:end]
}

func (b *BumpAllocator) Reset() {
	b.offset = 0
}

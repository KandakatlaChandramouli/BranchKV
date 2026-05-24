package ringbuffer

type RingBuffer struct {
	data [][]byte
	head int
	tail int
	size int
}

func NewRingBuffer(
	size int,
) *RingBuffer {

	return &RingBuffer{
		data: make(
			[][]byte,
			size,
		),
		size: size,
	}
}

func (r *RingBuffer) Push(
	v []byte,
) {

	r.data[r.tail] = v

	r.tail = (r.tail + 1) % r.size
}

func (r *RingBuffer) Pop() []byte {

	v := r.data[r.head]

	r.head = (r.head + 1) % r.size

	return v
}

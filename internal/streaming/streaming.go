package streaming

import (
	"sync"
)

type Stream struct {
	data [][]byte
	mu   sync.Mutex
}

func NewStream() *Stream {

	return &Stream{
		data: make([][]byte, 0),
	}
}

func (s *Stream) Publish(
	v []byte,
) {

	s.mu.Lock()
	defer s.mu.Unlock()

	s.data = append(
		s.data,
		v,
	)
}

func (s *Stream) Size() int {

	s.mu.Lock()
	defer s.mu.Unlock()

	return len(s.data)
}

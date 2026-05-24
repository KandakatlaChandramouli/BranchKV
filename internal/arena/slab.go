package arena

import (
	"sync"
)

type Slab struct {
	PageSize int
	Pages    [][]float32

	mu sync.Mutex
}

func NewSlab(
	pageSize int,
	capacity int,
) *Slab {

	s := &Slab{
		PageSize: pageSize,
		Pages:    make([][]float32, 0, capacity),
	}

	for i := 0; i < capacity; i++ {
		page := make([]float32, pageSize)

		s.Pages = append(s.Pages, page)
	}

	return s
}

func (s *Slab) Allocate() []float32 {

	s.mu.Lock()
	defer s.mu.Unlock()

	n := len(s.Pages)

	if n == 0 {
		return make([]float32, s.PageSize)
	}

	page := s.Pages[n-1]

	s.Pages = s.Pages[:n-1]

	return page
}

func (s *Slab) Free(page []float32) {

	s.mu.Lock()
	defer s.mu.Unlock()

	s.Pages = append(s.Pages, page)
}

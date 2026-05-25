package resident

import (
	"sync"
)

type Set struct {
	mu    sync.RWMutex
	pages map[uint64]struct{}
}

func NewSet() *Set {

	return &Set{
		pages: make(
			map[uint64]struct{},
		),
	}
}

func (s *Set) Add(
	page uint64,
) {

	s.mu.Lock()
	defer s.mu.Unlock()

	s.pages[page] = struct{}{}
}

func (s *Set) Remove(
	page uint64,
) {

	s.mu.Lock()
	defer s.mu.Unlock()

	delete(
		s.pages,
		page,
	)
}

func (s *Set) Contains(
	page uint64,
) bool {

	s.mu.RLock()
	defer s.mu.RUnlock()

	_, ok := s.pages[page]

	return ok
}

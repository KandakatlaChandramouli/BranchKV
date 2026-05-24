package multiversion

import (
	"sync"
)

type Version struct {
	ID uint64
}

type Store struct {
	versions []Version
	mu       sync.Mutex
}

func NewStore() *Store {

	return &Store{
		versions: make([]Version, 0),
	}
}

func (s *Store) Add(
	id uint64,
) {

	s.mu.Lock()
	defer s.mu.Unlock()

	s.versions = append(
		s.versions,
		Version{
			ID: id,
		},
	)
}

func (s *Store) Size() int {

	s.mu.Lock()
	defer s.mu.Unlock()

	return len(s.versions)
}

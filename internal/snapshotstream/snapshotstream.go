package snapshotstream

import (
	"sync"
)

type Snapshot struct {
	Version uint64
}

type Stream struct {
	snapshots []Snapshot
	mu        sync.Mutex
}

func NewStream() *Stream {

	return &Stream{
		snapshots: make([]Snapshot, 0),
	}
}

func (s *Stream) Push(
	version uint64,
) {

	s.mu.Lock()
	defer s.mu.Unlock()

	s.snapshots = append(
		s.snapshots,
		Snapshot{
			Version: version,
		},
	)
}

func (s *Stream) Size() int {

	s.mu.Lock()
	defer s.mu.Unlock()

	return len(s.snapshots)
}

package segmentmanager

import (
	"sync"
)

type Segment struct {
	ID uint64
}

type Manager struct {
	segments []Segment
	mu       sync.Mutex
}

func NewManager() *Manager {

	return &Manager{
		segments: make([]Segment, 0),
	}
}

func (m *Manager) Add(
	id uint64,
) {

	m.mu.Lock()
	defer m.mu.Unlock()

	m.segments = append(
		m.segments,
		Segment{
			ID: id,
		},
	)
}

func (m *Manager) Size() int {

	m.mu.Lock()
	defer m.mu.Unlock()

	return len(m.segments)
}

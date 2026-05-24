package checkpointing

import (
	"sync"
)

type Checkpoint struct {
	Version uint64
}

type Manager struct {
	checkpoints []Checkpoint
	mu          sync.Mutex
}

func NewManager() *Manager {

	return &Manager{
		checkpoints: make([]Checkpoint, 0),
	}
}

func (m *Manager) Save(
	version uint64,
) {

	m.mu.Lock()
	defer m.mu.Unlock()

	m.checkpoints = append(
		m.checkpoints,
		Checkpoint{
			Version: version,
		},
	)
}

func (m *Manager) Size() int {

	m.mu.Lock()
	defer m.mu.Unlock()

	return len(m.checkpoints)
}

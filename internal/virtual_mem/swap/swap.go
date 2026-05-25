package swap

import (
	"branchkv-core/internal/virtual_mem"
	"sync"
)

type Manager struct {
	mu    sync.RWMutex
	pages map[uint64][]float32
}

func NewManager() *Manager {

	return &Manager{
		pages: make(
			map[uint64][]float32,
		),
	}
}

func (m *Manager) SwapOut(
	frame *virtual_mem.PhysicalFrame,
) {

	m.mu.Lock()
	defer m.mu.Unlock()

	buf := make(
		[]float32,
		len(frame.Data),
	)

	copy(
		buf,
		frame.Data,
	)

	m.pages[frame.ID] = buf
}

func (m *Manager) SwapIn(
	frame *virtual_mem.PhysicalFrame,
) bool {

	m.mu.RLock()
	data, ok := m.pages[frame.ID]
	m.mu.RUnlock()

	if !ok {
		return false
	}

	copy(
		frame.Data,
		data,
	)

	return true
}

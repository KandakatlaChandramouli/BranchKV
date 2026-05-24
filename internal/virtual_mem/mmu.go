package virtual_mem

import (
	"branchkv-core/internal/arena"
	"fmt"
	"sync"
	"sync/atomic"
)

type MMU struct {
	mu         sync.RWMutex
	nextPageID atomic.Uint64
	virtualMap map[uint64]*VirtualDescriptor

	arena *arena.Arena
}

func NewMMU() *MMU {

	return &MMU{
		virtualMap: make(map[uint64]*VirtualDescriptor),
		arena:      arena.NewArena(),
	}
}

func (m *MMU) Allocate(
	logicalID uint64,
	size int,
) *VirtualDescriptor {

	pageID := m.nextPageID.Add(1)

	pageData := m.arena.Allocate(size)

	page := &PhysicalPage{
		ID:   pageID,
		Data: pageData[:size],
	}

	page.RefCount.Store(1)

	desc := &VirtualDescriptor{
		LogicalID: logicalID,
		Page:      page,
	}

	m.mu.Lock()
	m.virtualMap[logicalID] = desc
	m.mu.Unlock()

	return desc
}

func (m *MMU) Fork(
	srcLogicalID uint64,
	dstLogicalID uint64,
) error {

	m.mu.Lock()
	defer m.mu.Unlock()

	src, ok := m.virtualMap[srcLogicalID]

	if !ok {
		return fmt.Errorf("source page missing")
	}

	src.Page.IncRef()

	m.virtualMap[dstLogicalID] = &VirtualDescriptor{
		LogicalID: dstLogicalID,
		Page:      src.Page,
	}

	return nil
}

func (m *MMU) Resolve(
	logicalID uint64,
) (*VirtualDescriptor, bool) {

	m.mu.RLock()
	defer m.mu.RUnlock()

	d, ok := m.virtualMap[logicalID]

	return d, ok
}

package virtual_mem

import (
	"errors"
	"sync"
	"sync/atomic"
)

const (
	DefaultBlockSize = 256
	DefaultPoolSize  = 4096
)

var (
	ErrOutOfMemory = errors.New(
		"out of physical memory",
	)
)

type VirtualDescriptor struct {
	VirtualPageID  uint64
	PhysicalPageID uint64
	Offset         uint64
	Length         uint64
}

type PhysicalBlock struct {
	ID       uint64
	Data     []float32
	RefCount atomic.Int64
}

type PhysicalBlockAllocator struct {
	blocks []*PhysicalBlock
	free   []uint64
	mu     sync.Mutex
}

func NewPhysicalBlockAllocator(
	poolSize int,
	blockSize int,
) *PhysicalBlockAllocator {

	blocks := make(
		[]*PhysicalBlock,
		poolSize,
	)

	free := make(
		[]uint64,
		0,
		poolSize,
	)

	for i := 0; i < poolSize; i++ {

		blocks[i] = &PhysicalBlock{
			ID: uint64(i),
			Data: make(
				[]float32,
				blockSize,
			),
		}

		free = append(
			free,
			uint64(i),
		)
	}

	return &PhysicalBlockAllocator{
		blocks: blocks,
		free:   free,
	}
}

func (a *PhysicalBlockAllocator) Allocate() (
	*PhysicalBlock,
	error,
) {

	a.mu.Lock()
	defer a.mu.Unlock()

	if len(a.free) == 0 {

		return nil,
			ErrOutOfMemory
	}

	last := len(a.free) - 1

	id := a.free[last]

	a.free = a.free[:last]

	block := a.blocks[id]

	block.RefCount.Store(1)

	for i := range block.Data {
		block.Data[i] = 0
	}

	return block,
		nil
}

func (a *PhysicalBlockAllocator) Retain(
	block *PhysicalBlock,
) {

	block.RefCount.Add(1)
}

func (a *PhysicalBlockAllocator) Release(
	block *PhysicalBlock,
) {

	v := block.RefCount.Add(-1)

	if v > 0 {
		return
	}

	a.mu.Lock()
	defer a.mu.Unlock()

	a.free = append(
		a.free,
		block.ID,
	)
}

type VirtualPageTable struct {
	pages map[uint64]*PhysicalBlock
	mu    sync.RWMutex
}

func NewVirtualPageTable() *VirtualPageTable {

	return &VirtualPageTable{
		pages: make(
			map[uint64]*PhysicalBlock,
		),
	}
}

func (v *VirtualPageTable) Map(
	virtual uint64,
	block *PhysicalBlock,
) {

	v.mu.Lock()
	defer v.mu.Unlock()

	v.pages[virtual] = block
}

func (v *VirtualPageTable) Resolve(
	virtual uint64,
) (
	*PhysicalBlock,
	bool,
) {

	v.mu.RLock()
	defer v.mu.RUnlock()

	block, ok := v.pages[virtual]

	return block,
		ok
}

type MMU struct {
	allocator *PhysicalBlockAllocator
	table     *VirtualPageTable
	mu        sync.Mutex
}

func NewMMU() *MMU {

	return &MMU{
		allocator: NewPhysicalBlockAllocator(
			DefaultPoolSize,
			DefaultBlockSize,
		),
		table: NewVirtualPageTable(),
	}
}

func (m *MMU) AllocatePage(
	virtual uint64,
) error {

	block, err := m.allocator.Allocate()

	if err != nil {
		return err
	}

	m.table.Map(
		virtual,
		block,
	)

	return nil
}

func (m *MMU) AllocateDescriptor(
	virtual uint64,
	length uint64,
) (
	VirtualDescriptor,
	error,
) {

	err := m.AllocatePage(
		virtual,
	)

	if err != nil {

		return VirtualDescriptor{},
			err
	}

	block, ok := m.Resolve(
		virtual,
	)

	if !ok {

		return VirtualDescriptor{},
			errors.New(
				"descriptor allocation failed",
			)
	}

	return VirtualDescriptor{
		VirtualPageID:  virtual,
		PhysicalPageID: block.ID,
		Offset:         0,
		Length:         length,
	}, nil
}

func (m *MMU) Fork(
	parent uint64,
	child uint64,
) error {

	block, ok := m.Resolve(
		parent,
	)

	if !ok {

		return errors.New(
			"parent page missing",
		)
	}

	m.allocator.Retain(
		block,
	)

	m.table.Map(
		child,
		block,
	)

	return nil
}

func (m *MMU) Write(
	virtual uint64,
	index int,
	value float32,
) error {

	m.mu.Lock()
	defer m.mu.Unlock()

	block, ok := m.Resolve(
		virtual,
	)

	if !ok {

		return errors.New(
			"virtual page missing",
		)
	}

	if index < 0 ||
		index >= len(block.Data) {

		return errors.New(
			"index out of bounds",
		)
	}

	if block.RefCount.Load() > 1 {

		newBlock, err := m.allocator.Allocate()

		if err != nil {
			return err
		}

		copy(
			newBlock.Data,
			block.Data,
		)

		m.allocator.Release(
			block,
		)

		block = newBlock

		m.table.Map(
			virtual,
			block,
		)
	}

	block.Data[index] = value

	return nil
}

func (m *MMU) Read(
	virtual uint64,
	index int,
) (
	float32,
	error,
) {

	block, ok := m.Resolve(
		virtual,
	)

	if !ok {

		return 0,
			errors.New(
				"virtual page missing",
			)
	}

	if index < 0 ||
		index >= len(block.Data) {

		return 0,
			errors.New(
				"index out of bounds",
			)
	}

	return block.Data[index],
		nil
}

func (m *MMU) Resolve(
	virtual uint64,
) (
	*PhysicalBlock,
	bool,
) {

	return m.table.Resolve(
		virtual,
	)
}

func (m *MMU) RefCount(
	virtual uint64,
) int64 {

	block, ok := m.Resolve(
		virtual,
	)

	if !ok {
		return 0
	}

	return block.RefCount.Load()
}

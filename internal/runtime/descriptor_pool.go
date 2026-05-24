package runtime

import (
	"branchkv-core/internal/virtual_mem"
	"sync"
)

type DescriptorPool struct {
	descriptors []virtual_mem.VirtualDescriptor
	mu          sync.Mutex
}

func NewDescriptorPool() *DescriptorPool {

	return &DescriptorPool{
		descriptors: make(
			[]virtual_mem.VirtualDescriptor,
			0,
		),
	}
}

func (p *DescriptorPool) Add(
	desc virtual_mem.VirtualDescriptor,
) {

	p.mu.Lock()
	defer p.mu.Unlock()

	p.descriptors = append(
		p.descriptors,
		desc,
	)
}

func (p *DescriptorPool) Size() int {

	p.mu.Lock()
	defer p.mu.Unlock()

	return len(
		p.descriptors,
	)
}

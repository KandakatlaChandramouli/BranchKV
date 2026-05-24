package runtime

import (
	"branchkv-core/internal/virtual_mem"
	"sync"
)

var DescriptorPool = sync.Pool{
	New: func() any {
		return &virtual_mem.VirtualDescriptor{}
	},
}

func AcquireDescriptor() *virtual_mem.VirtualDescriptor {

	d := DescriptorPool.Get().(*virtual_mem.VirtualDescriptor)

	*d = virtual_mem.VirtualDescriptor{}

	return d
}

func ReleaseDescriptor(
	d *virtual_mem.VirtualDescriptor,
) {
	DescriptorPool.Put(d)
}

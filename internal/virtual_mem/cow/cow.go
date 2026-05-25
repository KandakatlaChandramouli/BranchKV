package cow

import "branchkv-core/internal/virtual_mem"

type CopyOnWrite struct {
	allocator *virtual_mem.FrameAllocator
	table     *virtual_mem.PageTable
}

func NewCopyOnWrite(
	allocator *virtual_mem.FrameAllocator,
	table *virtual_mem.PageTable,
) *CopyOnWrite {

	return &CopyOnWrite{
		allocator: allocator,
		table:     table,
	}
}

func (c *CopyOnWrite) Fork(
	virtual uint64,
	newVirtual uint64,
) bool {

	frame, ok := c.table.Resolve(
		virtual,
	)

	if !ok {
		return false
	}

	frame.Retain()

	c.table.Map(
		newVirtual,
		frame,
	)

	return true
}

func (c *CopyOnWrite) Write(
	virtual uint64,
	offset int,
	value float32,
) bool {

	frame, ok := c.table.Resolve(
		virtual,
	)

	if !ok {
		return false
	}

	if frame.RefCount.Load() > 1 {

		newFrame, ok := c.allocator.Alloc()

		if !ok {
			return false
		}

		copy(
			newFrame.Data,
			frame.Data,
		)

		frame.Release()

		c.table.Map(
			virtual,
			newFrame,
		)

		frame = newFrame
	}

	frame.Data[offset] = value

	frame.MarkDirty()

	return true
}

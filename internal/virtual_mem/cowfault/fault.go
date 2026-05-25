package cowfault

import (
	"branchkv-core/internal/virtual_mem"
	"branchkv-core/internal/virtual_mem/cow"
)

func Handle(
	frame *virtual_mem.PhysicalFrame,
) *virtual_mem.PhysicalFrame {

	if frame.RefCount.Load() <= 1 {
		return frame
	}

	frame.RefCount.Add(-1)

	return cow.Clone(
		frame,
	)
}

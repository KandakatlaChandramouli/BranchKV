package cow

import (
	"branchkv-core/internal/virtual_mem"
)

func Clone(
	frame *virtual_mem.PhysicalFrame,
) *virtual_mem.PhysicalFrame {

	next := virtual_mem.NewPhysicalFrame(
		frame.ID+1,
		len(
			frame.Data,
		),
	)

	copy(
		next.Data,
		frame.Data,
	)

	return next
}

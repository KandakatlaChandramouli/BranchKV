package translation

import "branchkv-core/internal/virtual_mem"

type Result struct {
	Frame  *virtual_mem.PhysicalFrame
	Offset uint64
	Hit    bool
}

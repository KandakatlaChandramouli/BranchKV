package conflict

import "branchkv-core/internal/virtual_mem/delta"

func Detect(
	left delta.Delta,
	right delta.Delta,
) bool {

	return left.PageID == right.PageID &&
		left.Offset == right.Offset
}

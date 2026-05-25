package rollback

import (
	"branchkv-core/internal/virtual_mem/snapshotspace"
)

func Restore(
	target *snapshotspace.Space,
	parent *snapshotspace.Space,
) {

	target.Root = parent.Root
}

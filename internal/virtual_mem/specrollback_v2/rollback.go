package specrollback_v2

import (
	"branchkv-core/internal/virtual_mem/specsnapshot"
)

func Restore(
	target *specsnapshot.Snapshot,
	source *specsnapshot.Snapshot,
) {

	target.Root = source.Root
}

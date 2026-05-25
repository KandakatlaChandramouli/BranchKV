package checkpointdelta

import "branchkv-core/internal/virtual_mem/delta"

type Snapshot struct {
	Entries []delta.Delta
}

func NewSnapshot(
	entries []delta.Delta,
) *Snapshot {

	return &Snapshot{
		Entries: entries,
	}
}

package snapshots

type Snapshot struct {
	ID         uint64
	ParentID   uint64
	Generation uint64
}

func NewSnapshot(
	id uint64,
	parent uint64,
	gen uint64,
) *Snapshot {

	return &Snapshot{
		ID:         id,
		ParentID:   parent,
		Generation: gen,
	}
}

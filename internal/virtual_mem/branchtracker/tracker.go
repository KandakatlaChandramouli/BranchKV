package branchtracker

type Tracker struct {
	active map[uint64]struct{}
}

func NewTracker() *Tracker {

	return &Tracker{
		active: make(
			map[uint64]struct{},
		),
	}
}

func (t *Tracker) Add(
	id uint64,
) {

	t.active[id] = struct{}{}
}

func (t *Tracker) Exists(
	id uint64,
) bool {

	_, ok := t.active[id]

	return ok
}

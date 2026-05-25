package eviction

type Evictor struct {
	pages []uint64
}

func NewEvictor() *Evictor {

	return &Evictor{
		pages: make(
			[]uint64,
			0,
		),
	}
}

func (e *Evictor) Add(
	page uint64,
) {

	e.pages = append(
		e.pages,
		page,
	)
}

func (e *Evictor) Evict() uint64 {

	if len(e.pages) == 0 {
		return 0
	}

	page := e.pages[0]

	e.pages = e.pages[1:]

	return page
}

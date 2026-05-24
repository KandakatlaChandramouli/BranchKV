package pager

type Page struct {
	ID   uint64
	Data []byte
}

type Pager struct {
	pages map[uint64]*Page
}

func NewPager() *Pager {

	return &Pager{
		pages: make(
			map[uint64]*Page,
		),
	}
}

func (p *Pager) Allocate(
	id uint64,
	size int,
) *Page {

	page := &Page{
		ID: id,
		Data: make(
			[]byte,
			size,
		),
	}

	p.pages[id] = page

	return page
}

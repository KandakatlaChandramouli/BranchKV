package compactor

type SSTable struct {
	ID uint64
}

type Compactor struct {
	tables []SSTable
}

func NewCompactor() *Compactor {

	return &Compactor{
		tables: make(
			[]SSTable,
			0,
		),
	}
}

func (c *Compactor) Add(
	id uint64,
) {

	c.tables = append(
		c.tables,
		SSTable{
			ID: id,
		},
	)
}

func (c *Compactor) Compact() int {

	if len(c.tables) <= 1 {
		return len(c.tables)
	}

	c.tables = c.tables[:1]

	return len(c.tables)
}

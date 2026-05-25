package writebatch

type Operation struct {
	Key   string
	Value []byte
}

type Batch struct {
	ops []Operation
}

func NewBatch() *Batch {

	return &Batch{
		ops: make(
			[]Operation,
			0,
		),
	}
}

func (b *Batch) Put(
	key string,
	value []byte,
) {

	b.ops = append(
		b.ops,
		Operation{
			Key:   key,
			Value: value,
		},
	)
}

func (b *Batch) Size() int {
	return len(b.ops)
}

package sstable

type Entry struct {
	Key   string
	Value []byte
}

type SSTable struct {
	Entries []Entry
}

func NewSSTable() *SSTable {

	return &SSTable{
		Entries: make(
			[]Entry,
			0,
		),
	}
}

func NewTable() *SSTable {

	return NewSSTable()
}

func (t *SSTable) Add(
	key string,
	value []byte,
) {

	t.Entries = append(
		t.Entries,
		Entry{
			Key:   key,
			Value: value,
		},
	)
}

func (t *SSTable) Size() int {
	return len(t.Entries)
}

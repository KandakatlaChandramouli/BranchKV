package sstable

type Entry struct {
	Key   string
	Value []byte
}

type SSTable struct {
	entries []Entry
}

func NewSSTable() *SSTable {

	return &SSTable{
		entries: make([]Entry, 0),
	}
}

func (s *SSTable) Append(
	key string,
	value []byte,
) {

	s.entries = append(
		s.entries,
		Entry{
			Key:   key,
			Value: value,
		},
	)
}

func (s *SSTable) Size() int {
	return len(s.entries)
}

package logstructured

import (
	"sync"
)

type Record struct {
	Key   string
	Value []byte
}

type LogStructuredStore struct {
	log []Record
	mu  sync.Mutex
}

func NewLogStructuredStore() *LogStructuredStore {

	return &LogStructuredStore{
		log: make([]Record, 0),
	}
}

func (l *LogStructuredStore) Append(
	r Record,
) {

	l.mu.Lock()
	defer l.mu.Unlock()

	l.log = append(
		l.log,
		r,
	)
}

func (l *LogStructuredStore) Size() int {

	l.mu.Lock()
	defer l.mu.Unlock()

	return len(l.log)
}

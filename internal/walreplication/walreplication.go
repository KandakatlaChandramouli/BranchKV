package walreplication

import (
	"sync"
)

type Entry struct {
	LSN uint64
}

type Replication struct {
	logs []Entry
	mu   sync.Mutex
}

func NewReplication() *Replication {

	return &Replication{
		logs: make([]Entry, 0),
	}
}

func (w *Replication) Append(
	lsn uint64,
) {

	w.mu.Lock()
	defer w.mu.Unlock()

	w.logs = append(
		w.logs,
		Entry{
			LSN: lsn,
		},
	)
}

func (w *Replication) Size() int {

	w.mu.Lock()
	defer w.mu.Unlock()

	return len(w.logs)
}

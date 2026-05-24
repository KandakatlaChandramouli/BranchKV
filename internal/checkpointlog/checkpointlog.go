package checkpointlog

import (
	"sync"
)

type Entry struct {
	Version uint64
}

type CheckpointLog struct {
	entries []Entry
	mu      sync.Mutex
}

func NewCheckpointLog() *CheckpointLog {

	return &CheckpointLog{
		entries: make([]Entry, 0),
	}
}

func (c *CheckpointLog) Append(
	version uint64,
) {

	c.mu.Lock()
	defer c.mu.Unlock()

	c.entries = append(
		c.entries,
		Entry{
			Version: version,
		},
	)
}

func (c *CheckpointLog) Size() int {

	c.mu.Lock()
	defer c.mu.Unlock()

	return len(c.entries)
}

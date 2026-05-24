package metrics

import (
	"sync/atomic"
)

type RuntimeMetrics struct {
	Forks        atomic.Uint64
	Allocations  atomic.Uint64
	TrieSearches atomic.Uint64
	Writes       atomic.Uint64
}

func NewMetrics() *RuntimeMetrics {
	return &RuntimeMetrics{}
}

func (m *RuntimeMetrics) RecordFork() {
	m.Forks.Add(1)
}

func (m *RuntimeMetrics) RecordAllocation() {
	m.Allocations.Add(1)
}

func (m *RuntimeMetrics) RecordTrieSearch() {
	m.TrieSearches.Add(1)
}

func (m *RuntimeMetrics) RecordWrite() {
	m.Writes.Add(1)
}

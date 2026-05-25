package lineage

import "sync"

type Lineage struct {
	mu       sync.RWMutex
	children map[uint64][]uint64
}

func NewLineage() *Lineage {

	return &Lineage{
		children: make(
			map[uint64][]uint64,
		),
	}
}

func (l *Lineage) Link(
	parent uint64,
	child uint64,
) {

	l.mu.Lock()
	defer l.mu.Unlock()

	l.children[parent] = append(
		l.children[parent],
		child,
	)
}

func (l *Lineage) Children(
	parent uint64,
) []uint64 {

	l.mu.RLock()
	defer l.mu.RUnlock()

	return l.children[parent]
}

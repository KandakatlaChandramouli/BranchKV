package cachecoherence

import "sync"

type Line struct {
	Address uint64
	Dirty   bool
}

type Directory struct {
	mu    sync.RWMutex
	lines map[uint64]Line
}

func NewDirectory() *Directory {

	return &Directory{
		lines: make(
			map[uint64]Line,
		),
	}
}

func (d *Directory) Store(
	address uint64,
) {

	d.mu.Lock()
	defer d.mu.Unlock()

	d.lines[address] = Line{
		Address: address,
		Dirty:   true,
	}
}

func (d *Directory) Resolve(
	address uint64,
) (
	Line,
	bool,
) {

	d.mu.RLock()
	defer d.mu.RUnlock()

	line, ok := d.lines[address]

	return line, ok
}

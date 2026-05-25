package mmu

import "sync"

type Context struct {
	ID uint64
}

type MMU struct {
	mu      sync.RWMutex
	current *Context
}

func NewMMU() *MMU {

	return &MMU{}
}

func (m *MMU) Switch(
	ctx *Context,
) {

	m.mu.Lock()
	defer m.mu.Unlock()

	m.current = ctx
}

func (m *MMU) Current() *Context {

	m.mu.RLock()
	defer m.mu.RUnlock()

	return m.current
}

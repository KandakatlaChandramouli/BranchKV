package connectionpool

import (
	"sync"
)

type Conn struct {
	ID uint64
}

type Pool struct {
	conns []Conn
	mu    sync.Mutex
}

func NewPool() *Pool {

	return &Pool{
		conns: make([]Conn, 0),
	}
}

func (p *Pool) Add(
	id uint64,
) {

	p.mu.Lock()
	defer p.mu.Unlock()

	p.conns = append(
		p.conns,
		Conn{
			ID: id,
		},
	)
}

func (p *Pool) Size() int {

	p.mu.Lock()
	defer p.mu.Unlock()

	return len(p.conns)
}

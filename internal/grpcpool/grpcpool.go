package grpcpool

import (
	"sync"
)

type Connection struct {
	Address string
}

type Pool struct {
	connections []Connection
	mu          sync.Mutex
}

func NewPool() *Pool {

	return &Pool{
		connections: make([]Connection, 0),
	}
}

func (p *Pool) Add(
	addr string,
) {

	p.mu.Lock()
	defer p.mu.Unlock()

	p.connections = append(
		p.connections,
		Connection{
			Address: addr,
		},
	)
}

func (p *Pool) Size() int {

	p.mu.Lock()
	defer p.mu.Unlock()

	return len(p.connections)
}

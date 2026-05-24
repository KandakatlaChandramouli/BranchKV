package objectpool_rt

import (
	"sync"
)

type Pool struct {
	pool sync.Pool
}

func NewPool() *Pool {

	return &Pool{
		pool: sync.Pool{
			New: func() any {
				return make(
					[]byte,
					1024,
				)
			},
		},
	}
}

func (p *Pool) Get() []byte {

	return p.pool.Get().([]byte)
}

func (p *Pool) Put(
	v []byte,
) {

	p.pool.Put(v)
}

package pool

import "sync"

type GenericPool struct {
	pool sync.Pool
}

func NewGenericPool(newFn func() any) *GenericPool {
	return &GenericPool{
		pool: sync.Pool{
			New: newFn,
		},
	}
}

func (p *GenericPool) Get() any {
	return p.pool.Get()
}

func (p *GenericPool) Put(v any) {
	p.pool.Put(v)
}

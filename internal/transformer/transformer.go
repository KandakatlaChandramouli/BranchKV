package transformer

import (
	"sync/atomic"
)

type Transformer struct {
	layers atomic.Uint64
}

func NewTransformer() *Transformer {

	return &Transformer{}
}

func (t *Transformer) Forward() {
	t.layers.Add(1)
}

func (t *Transformer) Layers() uint64 {
	return t.layers.Load()
}

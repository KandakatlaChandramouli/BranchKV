package generation

import "sync/atomic"

type Generator struct {
	next atomic.Uint64
}

func NewGenerator() *Generator {

	return &Generator{}
}

func (g *Generator) Next() uint64 {

	return g.next.Add(1)
}

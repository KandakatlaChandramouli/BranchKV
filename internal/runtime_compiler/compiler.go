package runtime_compiler

import (
	"sync/atomic"
)

type Compiler struct {
	compiled atomic.Uint64
}

func NewCompiler() *Compiler {

	return &Compiler{}
}

func (c *Compiler) Compile() {
	c.compiled.Add(1)
}

func (c *Compiler) Count() uint64 {
	return c.compiled.Load()
}

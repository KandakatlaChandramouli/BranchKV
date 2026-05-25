package specexecutor

import (
	"branchkv-core/internal/virtual_mem/dagnodearena"
)

type Executor struct {
	queue []*dagnodearena.Node
}

func NewExecutor() *Executor {

	return &Executor{}
}

func (e *Executor) Submit(
	node *dagnodearena.Node,
) {

	e.queue = append(
		e.queue,
		node,
	)
}

func (e *Executor) Size() int {

	return len(e.queue)
}

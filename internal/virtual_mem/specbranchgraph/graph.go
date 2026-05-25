package specbranchgraph

import (
	"branchkv-core/internal/virtual_mem/dagnodearena"
)

func Link(
	parent *dagnodearena.Node,
	child *dagnodearena.Node,
) {

	child.Parent = parent

	parent.Children = append(
		parent.Children,
		child,
	)
}

package specbranchgraph

import (
        "branchkv-core/internal/virtual_mem/dagnodearena"
)

func Link(
        parent *dagnodearena.Node,
        child *dagnodearena.Node,
) {
        child.Parent = parent

        if parent.Count >= uint8(len(parent.Children)) {
                return
        }

        parent.Children[parent.Count] = child
        parent.Count++
}

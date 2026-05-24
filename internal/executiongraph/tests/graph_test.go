package tests

import (
	"branchkv-core/internal/executiongraph"
	"testing"
)

func TestExecutionGraph(
	t *testing.T,
) {

	g := executiongraph.NewExecutionGraph()

	g.AddNode(1)

	g.AddEdge(1, 2)

	if g.Nodes() != 1 {
		t.Fatal("graph failed")
	}
}

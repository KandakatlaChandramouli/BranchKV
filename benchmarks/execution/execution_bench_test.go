package execution

import (
	"branchkv-core/internal/eventbus"
	"branchkv-core/internal/executiongraph"
	"testing"
)

func BenchmarkExecutionRuntime(
	b *testing.B,
) {

	graph := executiongraph.NewExecutionGraph()

	bus := eventbus.NewBus()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		graph.AddNode(
			uint64(i),
		)

		bus.Publish(
			"runtime",
		)
	}
}

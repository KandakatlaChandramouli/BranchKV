package tests

import (
	"testing"

	"branchkv-core/internal/virtual_mem/cachecoherence"
	"branchkv-core/internal/virtual_mem/consistencygraph"
	"branchkv-core/internal/virtual_mem/storebuffer"
)

func TestConsistencyLayer(
	t *testing.T,
) {

	dir := cachecoherence.NewDirectory()

	dir.Store(
		0x1000,
	)

	if _, ok := dir.Resolve(
		0x1000,
	); !ok {

		t.Fatal("directory failure")
	}

	graph := consistencygraph.NewGraph()

	graph.Link(
		1,
		2,
	)

	if len(
		graph.Children(
			1,
		),
	) == 0 {

		t.Fatal("graph failure")
	}

	buffer := storebuffer.NewBuffer()

	buffer.Push(
		1,
		100,
	)

	if buffer.Size() == 0 {

		t.Fatal("buffer failure")
	}
}

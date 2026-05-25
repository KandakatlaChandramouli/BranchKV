package main

import (
	"branchkv-core/internal/bloom"
	"branchkv-core/internal/compactor"
	"branchkv-core/internal/eviction"
	"branchkv-core/internal/lsm"
	"branchkv-core/internal/replayengine"
	"fmt"
)

func main() {

	tree := lsm.NewLSMTree()

	tree.Put(
		"runtime",
		[]byte("healthy"),
	)

	filter := bloom.NewBloom()

	filter.Add(
		"runtime",
	)

	comp := compactor.NewCompactor()

	comp.Add(1)
	comp.Add(2)

	evictor := eviction.NewEvictor()

	evictor.Add(42)

	replay := replayengine.NewReplayEngine()

	replay.Append(
		1,
		[]byte("wal"),
	)

	fmt.Println(
		"lsm:",
		tree.Size(),
	)

	fmt.Println(
		"compact:",
		comp.Compact(),
	)

	fmt.Println(
		"replay:",
		replay.Replay(),
	)

	fmt.Println(
		"evicted:",
		evictor.Evict(),
	)

	fmt.Println(
		"bloom:",
		filter.Contains("runtime"),
	)
}

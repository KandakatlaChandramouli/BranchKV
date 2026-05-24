package main

import (
	"branchkv-core/internal/buddyallocator"
	"branchkv-core/internal/epochgc"
	"branchkv-core/internal/lockfreequeue"
	"branchkv-core/internal/memoryarena"
	"branchkv-core/internal/numa"
	"branchkv-core/internal/objectpool_rt"
	"branchkv-core/internal/slaballocator"
	"fmt"
)

func main() {

	slab := slaballocator.NewAllocator()

	buf := slab.Allocate(128)

	buddy := buddyallocator.NewAllocator()

	buddy.Allocate(256)

	queue := lockfreequeue.NewQueue()

	queue.Push()

	ok := queue.Pop()

	pool := objectpool_rt.NewPool()

	obj := pool.Get()

	pool.Put(obj)

	gc := epochgc.NewGC()

	gc.Collect()

	arena := memoryarena.NewArena(
		2048,
	)

	mem := arena.Allocate(256)

	numaFabric := numa.NewNUMA()

	numaFabric.AddNode(0)

	numaFabric.AddNode(1)

	fmt.Println(
		"slab:",
		len(buf),
	)

	fmt.Println(
		"buddy:",
		buddy.Size(),
	)

	fmt.Println(
		"queue:",
		ok,
	)

	fmt.Println(
		"gc:",
		gc.Count(),
	)

	fmt.Println(
		"arena:",
		len(mem),
	)

	fmt.Println(
		"numa:",
		numaFabric.Size(),
	)
}

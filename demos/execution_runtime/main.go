package main

import (
	"branchkv-core/internal/allocator_rt"
	"branchkv-core/internal/checkpointlog"
	"branchkv-core/internal/dag"
	"branchkv-core/internal/eventbus"
	"branchkv-core/internal/executiongraph"
	"branchkv-core/internal/priorityqueue"
	"branchkv-core/internal/state_machine"
	"branchkv-core/internal/taskqueue"
	"container/heap"
	"fmt"
)

func main() {

	graph := executiongraph.NewExecutionGraph()

	graph.AddNode(1)

	graph.AddEdge(1, 2)

	d := dag.NewDAG()

	d.AddVertex(1)

	queue := taskqueue.NewQueue()

	queue.Push(
		taskqueue.Task{
			ID: 1,
		},
	)

	sm := state_machine.NewStateMachine()

	sm.Transition("running")

	log := checkpointlog.NewCheckpointLog()

	log.Append(1)

	bus := eventbus.NewBus()

	bus.Publish("runtime-start")

	pq := priorityqueue.NewPriorityQueue()

	heap.Push(
		pq,
		&priorityqueue.Item{
			Value:    "critical",
			Priority: 100,
		},
	)

	item := heap.Pop(pq).(*priorityqueue.Item)

	alloc := allocator_rt.NewRuntimeAllocator()

	alloc.Allocate()

	fmt.Println(
		"graph:",
		graph.Nodes(),
	)

	fmt.Println(
		"dag:",
		d.Size(),
	)

	fmt.Println(
		"queue:",
		queue.Size(),
	)

	fmt.Println(
		"state:",
		sm.State(),
	)

	fmt.Println(
		"log:",
		log.Size(),
	)

	fmt.Println(
		"bus:",
		bus.Size(),
	)

	fmt.Println(
		"priority:",
		item.Value,
	)

	fmt.Println(
		"alloc:",
		alloc.Count(),
	)
}

package main

import (
	"branchkv-core/internal/coroutines"
	"branchkv-core/internal/futures"
	"branchkv-core/internal/preemption"
	"branchkv-core/internal/runqueue"
	"branchkv-core/internal/workloadmodel"
	"branchkv-core/internal/workstealing"
	"fmt"
)

func main() {

	scheduler := workstealing.NewScheduler()

	scheduler.Push(1)

	scheduler.Push(2)

	ok := scheduler.Steal()

	queue := runqueue.NewQueue()

	queue.Push(1)

	v, _ := queue.Pop()

	preempt := preemption.NewEngine()

	preempt.Preempt()

	future := futures.NewFuture()

	go future.Complete(99)

	result := future.Await()

	coroutinesRuntime := coroutines.NewRuntime()

	coroutinesRuntime.Spawn()

	workload := workloadmodel.NewWorkload(
		"inference",
		50000,
	)

	fmt.Println(
		"steal:",
		ok,
	)

	fmt.Println(
		"queue:",
		v,
	)

	fmt.Println(
		"preempt:",
		preempt.Count(),
	)

	fmt.Println(
		"future:",
		result,
	)

	fmt.Println(
		"coroutines:",
		coroutinesRuntime.Count(),
	)

	fmt.Println(
		"workload:",
		workload.QPS,
	)
}

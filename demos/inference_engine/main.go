package main

import (
	"branchkv-core/internal/pipeline"
	"branchkv-core/internal/ringbuffer"
	"branchkv-core/internal/types"
	"branchkv-core/internal/workerpool"
	"fmt"
)

func main() {

	ring := ringbuffer.NewRingBuffer[[]types.TokenID](128)

	pool := workerpool.NewWorkerPool(
		4,
		32,
	)

	engine := pipeline.NewPipeline()

	input := []types.TokenID{
		1, 2, 3, 4, 5,
	}

	ring.Push(input)

	task, ok := ring.Pop()

	if !ok {
		panic("queue failed")
	}

	pool.Submit(func() {

		out := engine.Process(task)

		fmt.Println(
			"processed:",
			out,
		)
	})

	pool.Wait()
}

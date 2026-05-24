package main

import (
	"branchkv-core/internal/allreduce"
	"branchkv-core/internal/checkpointing"
	"branchkv-core/internal/cuda"
	"branchkv-core/internal/gpumemory"
	"branchkv-core/internal/nccl"
	"branchkv-core/internal/pipelineparallel"
	"branchkv-core/internal/state_sync"
	"branchkv-core/internal/streamexecutor"
	"branchkv-core/internal/tensorparallel"
	"fmt"
)

func main() {

	cudaRuntime := cuda.NewRuntime()

	cudaRuntime.Launch()

	ncclRuntime := nccl.NewNCCL()

	ncclRuntime.AllReduce()

	out := allreduce.Sum(
		[][]float32{
			{1, 2, 3},
			{4, 5, 6},
		},
	)

	tensor := tensorparallel.NewTensorParallel()

	tensor.AddShard(1)

	tensor.AddShard(2)

	pipe := pipelineparallel.NewPipeline()

	pipe.AddStage(1)

	pipe.AddStage(2)

	checkpoints := checkpointing.NewManager()

	checkpoints.Save(1)

	memory := gpumemory.NewPool()

	memory.Allocate(4096)

	streams := streamexecutor.NewExecutor()

	streams.Dispatch()

	syncer := state_sync.NewSync()

	syncer.Run()

	fmt.Println(
		"cuda:",
		cudaRuntime.Count(),
	)

	fmt.Println(
		"nccl:",
		ncclRuntime.Count(),
	)

	fmt.Println(
		"allreduce:",
		len(out),
	)

	fmt.Println(
		"tensor:",
		tensor.Size(),
	)

	fmt.Println(
		"pipeline:",
		pipe.Size(),
	)

	fmt.Println(
		"checkpoint:",
		checkpoints.Size(),
	)

	fmt.Println(
		"memory:",
		memory.Usage(),
	)

	fmt.Println(
		"streams:",
		streams.Count(),
	)

	fmt.Println(
		"sync:",
		syncer.Count(),
	)
}

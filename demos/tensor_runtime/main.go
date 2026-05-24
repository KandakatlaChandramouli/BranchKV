package main

import (
	"branchkv-core/internal/executionplanner"
	"branchkv-core/internal/gpu_scheduler"
	"branchkv-core/internal/kernelcache"
	"branchkv-core/internal/matmul"
	"branchkv-core/internal/operator_fusion"
	"branchkv-core/internal/runtime_compiler"
	"branchkv-core/internal/tensor"
	"branchkv-core/internal/tensorgraph"
	"fmt"
)

func main() {

	t := tensor.NewTensor(
		[]int{2, 2},
		[]float32{
			1, 2, 3, 4,
		},
	)

	out := matmul.Multiply(
		[][]float32{
			{1, 2},
		},
		[][]float32{
			{3},
			{4},
		},
	)

	scheduler := gpu_scheduler.NewScheduler()

	scheduler.Dispatch()

	fusion := operator_fusion.NewFusionEngine()

	fusion.Fuse("matmul+gelu")

	cache := kernelcache.NewCache()

	cache.Store("matmul")

	planner := executionplanner.NewPlanner()

	planner.Add("dispatch-matmul")

	graph := tensorgraph.NewTensorGraph()

	graph.AddNode(1)

	compiler := runtime_compiler.NewCompiler()

	compiler.Compile()

	fmt.Println(
		"tensor:",
		t.Size(),
	)

	fmt.Println(
		"matmul:",
		out[0][0],
	)

	fmt.Println(
		"dispatch:",
		scheduler.Count(),
	)

	fmt.Println(
		"fusion:",
		fusion.Size(),
	)

	fmt.Println(
		"cache:",
		cache.Size(),
	)

	fmt.Println(
		"planner:",
		planner.Size(),
	)

	fmt.Println(
		"graph:",
		graph.Size(),
	)

	fmt.Println(
		"compiled:",
		compiler.Count(),
	)
}

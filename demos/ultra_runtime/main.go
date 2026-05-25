package main

import (
	"branchkv-core/internal/runtimeepoch"
	"branchkv-core/internal/runtimefabric"
	"branchkv-core/internal/runtimegraph"
	"branchkv-core/internal/runtimepipeline"
	"branchkv-core/internal/runtimevector"
	"fmt"
)

func main() {

	fabric := runtimefabric.NewRuntime()
	vector := runtimevector.NewRuntime()
	graph := runtimegraph.NewRuntime()
	pipe := runtimepipeline.NewRuntime()
	epoch := runtimeepoch.NewRuntime()

	fabric.Add()
	vector.Add()
	graph.Add()
	pipe.Add()
	epoch.Add()

	fmt.Println("fabric:", fabric.Size())
	fmt.Println("vector:", vector.Size())
	fmt.Println("graph:", graph.Size())
	fmt.Println("pipeline:", pipe.Size())
	fmt.Println("epoch:", epoch.Size())
}

package main

import (
	"branchkv-core/internal/runtimeanalytics"
	"branchkv-core/internal/runtimefabric_v2"
	"branchkv-core/internal/runtimegpu"
	"branchkv-core/internal/runtimeinference"
	"branchkv-core/internal/runtimeplanner"
	"fmt"
)

func main() {

	fabric := runtimefabric_v2.NewRuntime()
	planner := runtimeplanner.NewRuntime()
	analytics := runtimeanalytics.NewRuntime()
	gpu := runtimegpu.NewRuntime()
	inference := runtimeinference.NewRuntime()

	fabric.Add()
	planner.Add()
	analytics.Add()
	gpu.Add()
	inference.Add()

	fmt.Println("fabric:", fabric.Size())
	fmt.Println("planner:", planner.Size())
	fmt.Println("analytics:", analytics.Size())
	fmt.Println("gpu:", gpu.Size())
	fmt.Println("inference:", inference.Size())
}

package main

import (
	"branchkv-core/internal/runtimecluster"
	"branchkv-core/internal/runtimeexecutor"
	"branchkv-core/internal/runtimegateway"
	"branchkv-core/internal/runtimeorchestrator"
	"branchkv-core/internal/runtimewatchdog"
	"fmt"
)

func main() {

	cluster := runtimecluster.NewRuntime()
	orchestrator := runtimeorchestrator.NewRuntime()
	executor := runtimeexecutor.NewRuntime()
	gateway := runtimegateway.NewRuntime()
	watchdog := runtimewatchdog.NewRuntime()

	cluster.Add()
	orchestrator.Add()
	executor.Add()
	gateway.Add()
	watchdog.Add()

	fmt.Println("cluster:", cluster.Size())
	fmt.Println("orchestrator:", orchestrator.Size())
	fmt.Println("executor:", executor.Size())
	fmt.Println("gateway:", gateway.Size())
	fmt.Println("watchdog:", watchdog.Size())
}

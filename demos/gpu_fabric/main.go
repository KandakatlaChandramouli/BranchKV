package main

import (
	"branchkv-core/internal/consensus"
	"branchkv-core/internal/fabric"
	"branchkv-core/internal/gpu"
	"branchkv-core/internal/nvlink"
	"branchkv-core/internal/rdma"
	"branchkv-core/internal/telemetry"
	"fmt"
)

func main() {

	rdmaEngine := rdma.NewRDMAEngine()

	region := rdmaEngine.Allocate(
		16384,
	)

	gpu0 := gpu.NewGPUExecutor(0)

	gpu1 := gpu.NewGPUExecutor(1)

	link := nvlink.NewNVLinkFabric()

	link.Connect(0, 1)

	fabricEngine := fabric.NewFabric()

	fabricEngine.AddNode(1)
	fabricEngine.AddNode(2)

	tel := telemetry.NewTelemetry()

	tel.Record()

	cons := consensus.NewConsensus()

	cons.Submit(
		consensus.Vote{
			NodeID: 1,
			Value:  true,
		},
	)

	out := gpu0.Execute(
		[]float32{
			1, 2, 3, 4,
		},
	)

	_ = gpu1

	fmt.Println(
		"rdma region:",
		region.ID,
	)

	fmt.Println(
		"nvlink:",
		link.Size(),
	)

	fmt.Println(
		"fabric:",
		fabricEngine.Size(),
	)

	fmt.Println(
		"telemetry:",
		tel.Events.Load(),
	)

	fmt.Println(
		"votes:",
		cons.Count(),
	)

	fmt.Println(
		"gpu out:",
		len(out),
	)
}

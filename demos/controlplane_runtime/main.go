package main

import (
	"branchkv-core/internal/autoscaler"
	"branchkv-core/internal/backpressure"
	"branchkv-core/internal/controlplane"
	"branchkv-core/internal/failover"
	"branchkv-core/internal/heartbeat"
	"branchkv-core/internal/servicediscovery"
	"branchkv-core/internal/streaming"
	"branchkv-core/internal/topology"
	"fmt"
)

func main() {

	cp := controlplane.NewControlPlane()

	cp.Register(
		controlplane.RuntimeNode{
			ID:     1,
			Status: "healthy",
		},
	)

	discovery := servicediscovery.NewDiscovery()

	discovery.Register(
		servicediscovery.Service{
			Name: "branchkv-runtime",
			Host: "127.0.0.1",
		},
	)

	scaler := autoscaler.NewAutoScaler()

	scaler.ScaleUp()

	hb := heartbeat.NewHeartbeat()

	hb.Beat()

	top := topology.NewTopology()

	top.Connect(1, 2)

	fail := failover.NewFailover()

	fail.Trigger()

	stream := streaming.NewStream()

	stream.Publish(
		[]byte("runtime-state"),
	)

	bp := backpressure.NewBackpressure()

	bp.Drop()

	fmt.Println(
		"nodes:",
		cp.Size(),
	)

	fmt.Println(
		"autoscaled:",
		scaler.Nodes(),
	)

	fmt.Println(
		"heartbeat:",
		hb.Count(),
	)

	fmt.Println(
		"topology:",
		top.Size(),
	)

	fmt.Println(
		"failover:",
		fail.Count(),
	)

	fmt.Println(
		"stream:",
		stream.Size(),
	)

	fmt.Println(
		"backpressure:",
		bp.Count(),
	)
}

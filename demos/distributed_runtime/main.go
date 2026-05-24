package main

import (
	"branchkv-core/internal/cluster"
	"branchkv-core/internal/loadbalancer"
	"branchkv-core/internal/network"
	"branchkv-core/internal/node"
	"branchkv-core/internal/replication"
	"branchkv-core/internal/router"
	"fmt"
)

func main() {

	clusterEngine := cluster.NewCluster()

	for i := 0; i < 4; i++ {

		clusterEngine.AddNode(
			cluster.Node{
				ID:      uint64(i),
				Address: "localhost",
			},
		)
	}

	routerEngine := router.NewRouter()

	balancer := loadbalancer.NewLoadBalancer()

	replicator := replication.NewReplicator()

	nodes := make(
		[]*node.RuntimeNode,
		0,
	)

	for i := 0; i < 4; i++ {

		nodes = append(
			nodes,
			node.NewRuntimeNode(
				uint64(i),
			),
		)
	}

	key := "token-branch"

	target := routerEngine.Route(
		key,
		4,
	)

	balance := balancer.Next(4)

	msg := network.Message{
		NodeID:  uint64(target),
		Payload: []byte("branch state"),
	}

	nodes[target].Receive(msg)

	replicator.Replicate(msg)

	fmt.Println(
		"cluster size:",
		clusterEngine.Size(),
	)

	fmt.Println(
		"routed node:",
		target,
	)

	fmt.Println(
		"balanced node:",
		balance,
	)

	fmt.Println(
		"inbox:",
		nodes[target].InboxSize(),
	)
}

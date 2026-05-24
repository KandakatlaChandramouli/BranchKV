package main

import (
	"branchkv-core/internal/connectionpool"
	"branchkv-core/internal/grpcpool"
	"branchkv-core/internal/loadbalancer"
	"branchkv-core/internal/ratelimiter"
	"branchkv-core/internal/ringbuffer"
	"branchkv-core/internal/rpc"
	"branchkv-core/internal/servicemesh"
	"branchkv-core/internal/zerocopy"
	"fmt"
	"time"
)

func main() {

	go rpc.StartServer(
		":8088",
	)

	time.Sleep(
		time.Millisecond * 200,
	)

	out, err := rpc.Call(
		":8088",
		10,
	)

	if err != nil {
		panic(err)
	}

	grpc := grpcpool.NewPool()

	grpc.Add("localhost")

	balancer := loadbalancer.NewBalancer()

	node := balancer.Next(4)

	limiter := ratelimiter.NewLimiter(10)

	allowed := limiter.Allow()

	mesh := servicemesh.NewMesh()

	mesh.Connect(
		"gateway",
		"inference",
	)

	conns := connectionpool.NewPool()

	conns.Add(1)

	ring := ringbuffer.NewRingBuffer(8)

	ring.Push(
		[]byte("runtime"),
	)

	buf := ring.Pop()

	zero := zerocopy.Slice(
		buf,
		0,
		4,
	)

	fmt.Println(
		"rpc:",
		out,
	)

	fmt.Println(
		"grpc:",
		grpc.Size(),
	)

	fmt.Println(
		"node:",
		node,
	)

	fmt.Println(
		"allowed:",
		allowed,
	)

	fmt.Println(
		"mesh:",
		mesh.Size(),
	)

	fmt.Println(
		"conns:",
		conns.Size(),
	)

	fmt.Println(
		"zero:",
		string(zero),
	)
}

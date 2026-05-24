package main

import (
	"branchkv-core/internal/ringbuffer"
	"fmt"
)

func main() {

	ring := ringbuffer.NewRingBuffer(8)

	ring.Push(
		[]byte("runtime"),
	)

	out := ring.Pop()

	fmt.Println(
		string(out),
	)
}

package network

import (
	"branchkv-core/internal/ringbuffer"
	"testing"
)

func BenchmarkNetworkRuntime(
	b *testing.B,
) {

	r := ringbuffer.NewRingBuffer(1024)

	payload := []byte("runtime")

	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		r.Push(payload)

		_ = r.Pop()
	}
}

package benchmarks

import (
	"branchkv-core/internal/ringbuffer"
	"testing"
)

func BenchmarkRingBuffer(
	b *testing.B,
) {

	r := ringbuffer.NewRingBuffer[int](1024)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		r.Push(i)

		_, _ = r.Pop()
	}
}

package scheduler

import (
	"branchkv-core/internal/workstealing"
	"testing"
)

func BenchmarkScheduler(
	b *testing.B,
) {

	s := workstealing.NewScheduler()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		s.Push(uint64(i))

		s.Steal()
	}
}

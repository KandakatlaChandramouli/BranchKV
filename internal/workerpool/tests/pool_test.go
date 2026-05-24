package tests

import (
	"branchkv-core/internal/workerpool"
	"sync/atomic"
	"testing"
)

func TestWorkerPool(
	t *testing.T,
) {

	pool := workerpool.NewWorkerPool(
		4,
		16,
	)

	var counter atomic.Uint64

	for i := 0; i < 100; i++ {

		pool.Submit(func() {
			counter.Add(1)
		})
	}

	pool.Wait()

	if counter.Load() != 100 {
		t.Fatal("worker pool failed")
	}
}

package tests

import (
	"branchkv-core/internal/lockfreequeue"
	"testing"
)

func TestLockFreeQueue(
	t *testing.T,
) {

	q := lockfreequeue.NewQueue()

	q.Push()

	if !q.Pop() {
		t.Fatal("queue failed")
	}
}

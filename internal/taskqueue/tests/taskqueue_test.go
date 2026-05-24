package tests

import (
	"branchkv-core/internal/taskqueue"
	"testing"
)

func TestTaskQueue(
	t *testing.T,
) {

	q := taskqueue.NewQueue()

	q.Push(
		taskqueue.Task{
			ID: 1,
		},
	)

	if q.Size() != 1 {
		t.Fatal("queue failed")
	}
}

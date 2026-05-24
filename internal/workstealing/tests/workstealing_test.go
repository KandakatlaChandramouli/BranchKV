package tests

import (
	"branchkv-core/internal/workstealing"
	"testing"
)

func TestWorkStealing(
	t *testing.T,
) {

	s := workstealing.NewScheduler()

	s.Push(1)

	ok := s.Steal()

	if !ok {
		t.Fatal("steal failed")
	}
}

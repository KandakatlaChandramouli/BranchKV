package tests

import (
	"branchkv-core/internal/ringbuffer"
	"testing"
)

func TestRingBuffer(
	t *testing.T,
) {

	r := ringbuffer.NewRingBuffer[int](16)

	r.Push(42)

	v, ok := r.Pop()

	if !ok {
		t.Fatal("pop failed")
	}

	if v != 42 {
		t.Fatal("invalid value")
	}
}

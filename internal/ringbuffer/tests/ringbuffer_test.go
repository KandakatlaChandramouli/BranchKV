package tests

import (
	"branchkv-core/internal/ringbuffer"
	"testing"
)

func TestRingBuffer(
	t *testing.T,
) {

	r := ringbuffer.NewRingBuffer(4)

	r.Push([]byte("a"))

	out := r.Pop()

	if string(out) != "a" {
		t.Fatal("ring failed")
	}
}

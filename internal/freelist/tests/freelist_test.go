package tests

import (
	"branchkv-core/internal/freelist"
	"testing"
	"unsafe"
)

func TestFreeList(t *testing.T) {

	f := freelist.NewFreeList()

	v := 42

	f.Push(
		unsafe.Pointer(&v),
	)

	ptr := f.Pop()

	if ptr == nil {
		t.Fatal("pop failed")
	}
}

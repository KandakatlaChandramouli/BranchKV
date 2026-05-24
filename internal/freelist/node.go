package freelist

import (
	"sync/atomic"
	"unsafe"
)

type Node struct {
	Value unsafe.Pointer
	Next  unsafe.Pointer
}

type FreeList struct {
	Head atomic.Pointer[Node]
}

func NewFreeList() *FreeList {
	return &FreeList{}
}

func (f *FreeList) Push(ptr unsafe.Pointer) {

	node := &Node{
		Value: ptr,
	}

	for {

		head := f.Head.Load()

		node.Next = unsafe.Pointer(head)

		if f.Head.CompareAndSwap(
			head,
			node,
		) {
			return
		}
	}
}

func (f *FreeList) Pop() unsafe.Pointer {

	for {

		head := f.Head.Load()

		if head == nil {
			return nil
		}

		next := (*Node)(head.Next)

		if f.Head.CompareAndSwap(
			head,
			next,
		) {

			return head.Value
		}
	}
}

package priorityqueue

import (
	"container/heap"
)

type Item struct {
	Value    string
	Priority int
	Index    int
}

type PriorityQueue []*Item

func (p PriorityQueue) Len() int {
	return len(p)
}

func (p PriorityQueue) Less(
	i int,
	j int,
) bool {

	return p[i].Priority >
		p[j].Priority
}

func (p PriorityQueue) Swap(
	i int,
	j int,
) {

	p[i], p[j] = p[j], p[i]

	p[i].Index = i
	p[j].Index = j
}

func (p *PriorityQueue) Push(
	x any,
) {

	item := x.(*Item)

	item.Index = len(*p)

	*p = append(*p, item)
}

func (p *PriorityQueue) Pop() any {

	old := *p

	n := len(old)

	item := old[n-1]

	old[n-1] = nil

	item.Index = -1

	*p = old[:n-1]

	return item
}

func NewPriorityQueue() *PriorityQueue {

	pq := &PriorityQueue{}

	heap.Init(pq)

	return pq
}

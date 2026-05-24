package node

import (
	"branchkv-core/internal/network"
	"sync"
)

type RuntimeNode struct {
	ID uint64

	inbox []network.Message

	mu sync.Mutex
}

func NewRuntimeNode(
	id uint64,
) *RuntimeNode {

	return &RuntimeNode{
		ID: id,
		inbox: make(
			[]network.Message,
			0,
		),
	}
}

func (r *RuntimeNode) Receive(
	msg network.Message,
) {

	r.mu.Lock()
	defer r.mu.Unlock()

	r.inbox = append(
		r.inbox,
		msg,
	)
}

func (r *RuntimeNode) InboxSize() int {

	r.mu.Lock()
	defer r.mu.Unlock()

	return len(r.inbox)
}

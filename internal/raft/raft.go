package raft

import (
	"sync/atomic"
)

type RaftNode struct {
	Term atomic.Uint64
}

func NewRaftNode() *RaftNode {
	return &RaftNode{}
}

func (r *RaftNode) NextTerm() uint64 {
	return r.Term.Add(1)
}

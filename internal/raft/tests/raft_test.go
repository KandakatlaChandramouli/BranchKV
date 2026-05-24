package tests

import (
	"branchkv-core/internal/raft"
	"testing"
)

func TestRaft(
	t *testing.T,
) {

	node := raft.NewRaftNode()

	term := node.NextTerm()

	if term != 1 {
		t.Fatal("raft failed")
	}
}

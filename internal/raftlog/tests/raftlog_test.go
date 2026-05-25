package tests

import (
	"branchkv-core/internal/raftlog"
	"testing"
)

func TestRaftLog(
	t *testing.T,
) {

	log := raftlog.NewRaftLog()

	log.Append(
		1,
		[]byte("wal"),
	)

	if log.Size() != 1 {
		t.Fatal("raft append failed")
	}
}

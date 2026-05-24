package tests

import (
	"branchkv-core/internal/network"
	"branchkv-core/internal/replication"
	"testing"
)

func TestReplication(
	t *testing.T,
) {

	r := replication.NewReplicator()

	r.Replicate(
		network.Message{
			NodeID:  1,
			Payload: []byte("hello"),
		},
	)
}

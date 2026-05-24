package tests

import (
	"branchkv-core/internal/consensus"
	"testing"
)

func TestConsensus(
	t *testing.T,
) {

	c := consensus.NewConsensus()

	c.Submit(
		consensus.Vote{
			NodeID: 1,
			Value:  true,
		},
	)

	if c.Count() != 1 {
		t.Fatal("consensus failed")
	}
}

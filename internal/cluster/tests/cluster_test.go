package tests

import (
	"branchkv-core/internal/cluster"
	"testing"
)

func TestCluster(
	t *testing.T,
) {

	c := cluster.NewCluster()

	c.AddNode(
		cluster.Node{
			ID:      1,
			Address: "localhost",
		},
	)

	if c.Size() != 1 {
		t.Fatal("cluster failed")
	}
}

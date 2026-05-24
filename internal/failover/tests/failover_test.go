package tests

import (
	"branchkv-core/internal/failover"
	"testing"
)

func TestFailover(
	t *testing.T,
) {

	f := failover.NewFailover()

	f.Trigger()

	if f.Count() != 1 {
		t.Fatal("failover failed")
	}
}

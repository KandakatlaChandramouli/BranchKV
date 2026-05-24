package tests

import (
	"branchkv-core/internal/loadbalancer"
	"testing"
)

func TestLoadBalancer(
	t *testing.T,
) {

	b := loadbalancer.NewBalancer()

	v := b.Next(4)

	if v > 3 {
		t.Fatal("invalid node")
	}
}

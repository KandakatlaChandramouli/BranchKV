package tests

import (
	"branchkv-core/internal/loadbalancer"
	"testing"
)

func TestBalancer(
	t *testing.T,
) {

	l := loadbalancer.NewLoadBalancer()

	v := l.Next(4)

	if v >= 4 {
		t.Fatal("invalid balance")
	}
}

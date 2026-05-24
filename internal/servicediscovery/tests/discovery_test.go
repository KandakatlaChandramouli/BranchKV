package tests

import (
	"branchkv-core/internal/servicediscovery"
	"testing"
)

func TestDiscovery(
	t *testing.T,
) {

	d := servicediscovery.NewDiscovery()

	d.Register(
		servicediscovery.Service{
			Name: "runtime",
			Host: "localhost",
		},
	)

	_, ok := d.Resolve("runtime")

	if !ok {
		t.Fatal("resolve failed")
	}
}

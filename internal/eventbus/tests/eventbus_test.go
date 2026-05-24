package tests

import (
	"branchkv-core/internal/eventbus"
	"testing"
)

func TestEventBus(
	t *testing.T,
) {

	b := eventbus.NewBus()

	b.Publish("runtime")

	if b.Size() != 1 {
		t.Fatal("event failed")
	}
}

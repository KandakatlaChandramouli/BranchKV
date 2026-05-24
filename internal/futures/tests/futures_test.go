package tests

import (
	"branchkv-core/internal/futures"
	"testing"
)

func TestFuture(
	t *testing.T,
) {

	f := futures.NewFuture()

	go f.Complete(42)

	v := f.Await()

	if v != 42 {
		t.Fatal("future failed")
	}
}

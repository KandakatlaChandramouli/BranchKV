package tests

import (
	"branchkv-core/internal/trace"
	"testing"
)

func TestTrace(
	t *testing.T,
) {

	tr := trace.NewTracer()

	tr.Start("runtime")

	if tr.Size() != 1 {
		t.Fatal("trace failed")
	}
}

package tests

import (
	"branchkv-core/internal/threading"
	"testing"
)

func TestThread(
	t *testing.T,
) {

	tm := threading.NewThreadManager()

	th := tm.Spawn()

	if th.TID == 0 {
		t.Fatal("thread failed")
	}
}

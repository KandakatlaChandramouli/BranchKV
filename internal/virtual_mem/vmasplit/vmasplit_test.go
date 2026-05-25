package vmasplit

import "testing"

func TestRuntime(
	t *testing.T,
) {
	rt := NewRuntime()

	if rt.Next() == 0 {
		t.Fatal("invalid runtime increment")
	}
}

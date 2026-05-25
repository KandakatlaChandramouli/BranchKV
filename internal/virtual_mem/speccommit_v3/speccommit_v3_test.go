package speccommit_v3

import "testing"

func TestRuntime(
	t *testing.T,
) {
	rt := NewRuntime()

	if rt.Advance() == 0 {
		t.Fatal("advance failed")
	}
}

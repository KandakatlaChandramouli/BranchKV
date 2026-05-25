package stresspressure

import "testing"

func TestRuntime(
	t *testing.T,
) {

	rt := NewRuntime()

	if rt.Advance() == 0 {

		t.Fatal("runtime failure")
	}
}

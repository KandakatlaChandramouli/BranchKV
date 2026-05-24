package tests

import (
	"branchkv-core/internal/process"
	"testing"
)

func TestProcess(
	t *testing.T,
) {

	pm := process.NewProcessManager()

	p := pm.Spawn()

	if p.PID == 0 {
		t.Fatal("process failed")
	}
}

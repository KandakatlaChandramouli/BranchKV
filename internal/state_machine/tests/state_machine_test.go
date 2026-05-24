package tests

import (
	"branchkv-core/internal/state_machine"
	"testing"
)

func TestStateMachine(
	t *testing.T,
) {

	s := state_machine.NewStateMachine()

	s.Transition("running")

	if s.State() != "running" {
		t.Fatal("state failed")
	}
}

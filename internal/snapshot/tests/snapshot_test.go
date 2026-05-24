package tests

import (
	"branchkv-core/internal/snapshot"
	"testing"
)

type TestState struct {
	Value int
}

func TestSnapshot(
	t *testing.T,
) {

	engine := snapshot.NewSnapshotEngine(
		"snapshot.gob",
	)

	state := TestState{
		Value: 99,
	}

	err := engine.Save(state)

	if err != nil {
		t.Fatal(err)
	}

	var loaded TestState

	err = engine.Load(&loaded)

	if err != nil {
		t.Fatal(err)
	}

	if loaded.Value != 99 {
		t.Fatal("snapshot mismatch")
	}
}

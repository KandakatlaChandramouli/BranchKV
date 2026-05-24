package tests

import (
	"branchkv-core/internal/wal"
	"testing"
)

func TestWALAppend(
	t *testing.T,
) {

	w, err := wal.NewWAL(
		"test.wal",
	)

	if err != nil {
		t.Fatal(err)
	}

	defer w.Close()

	err = w.Append(42)

	if err != nil {
		t.Fatal(err)
	}
}

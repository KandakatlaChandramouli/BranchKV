package tests

import (
	"branchkv-core/internal/wal"
	"testing"
)

func TestWALAppend(
	t *testing.T,
) {

	w, err := wal.Open(
		"test.wal",
	)

	if err != nil {
		t.Fatal(err)
	}

	defer w.Close()

	err = w.Append(
		[]byte("branchkv"),
	)

	if err != nil {
		t.Fatal(err)
	}
}

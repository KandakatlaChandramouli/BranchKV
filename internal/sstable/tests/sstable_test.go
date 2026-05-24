package tests

import (
	"branchkv-core/internal/sstable"
	"testing"
)

func TestSSTable(
	t *testing.T,
) {

	table := sstable.NewSSTable()

	table.Add(
		"branchkv",
		[]byte("runtime"),
	)

	if table.Size() != 1 {
		t.Fatal("sstable insert failed")
	}
}

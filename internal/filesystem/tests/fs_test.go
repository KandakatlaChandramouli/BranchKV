package tests

import (
	"branchkv-core/internal/filesystem"
	"testing"
)

func TestFileSystem(
	t *testing.T,
) {

	fs := filesystem.NewFileSystem()

	fs.Create(
		"runtime",
		[]byte("branchkv"),
	)

	_, ok := fs.Read("runtime")

	if !ok {
		t.Fatal("read failed")
	}
}

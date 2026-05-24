package tests

import (
	"branchkv-core/internal/mmapengine"
	"testing"
)

func TestMMap(
	t *testing.T,
) {

	r, err := mmapengine.MapFile(
		"test.mmap",
		4096,
	)

	if err != nil {
		t.Fatal(err)
	}

	defer r.Unmap()

	r.Data[0] = 42

	if r.Data[0] != 42 {
		t.Fatal("mmap failed")
	}
}

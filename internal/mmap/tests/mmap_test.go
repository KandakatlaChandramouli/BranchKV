package tests

import (
	"branchkv-core/internal/mmap"
	"testing"
)

func TestMMapRegion(
	t *testing.T,
) {

	region, err := mmap.NewMMapRegion(
		"test.mmap",
		4096,
	)

	if err != nil {
		t.Fatal(err)
	}

	defer region.Close()

	if len(region.Data) != 4096 {
		t.Fatal("invalid mmap size")
	}
}

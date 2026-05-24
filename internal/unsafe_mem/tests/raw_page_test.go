package tests

import (
	"branchkv-core/internal/unsafe_mem"
	"testing"
)

func TestRawPageWriteRead(t *testing.T) {

	page := unsafe_mem.NewRawPage(1024)

	page.WriteFloat32(0, 42.5)

	val := page.ReadFloat32(0)

	if val != 42.5 {
		t.Fatal("unsafe memory read/write failed")
	}
}

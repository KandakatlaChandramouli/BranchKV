package tests

import (
	"branchkv-core/internal/bloomfilter"
	"testing"
)

func TestBloom(
	t *testing.T,
) {

	f := bloomfilter.NewFilter(1024)

	f.Add("runtime")

	if !f.Contains("runtime") {
		t.Fatal("bloom failed")
	}
}

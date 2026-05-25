package tests

import (
	"branchkv-core/internal/bloom"
	"testing"
)

func TestBloom(
	t *testing.T,
) {

	b := bloom.NewBloom()

	b.Add(
		"kv",
	)

	if !b.Contains("kv") {
		t.Fatal("bloom failed")
	}
}

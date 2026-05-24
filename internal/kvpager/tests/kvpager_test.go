package tests

import (
	"branchkv-core/internal/kvpager"
	"testing"
)

func TestKVPager(
	t *testing.T,
) {

	p := kvpager.NewKVPager()

	p.Allocate(
		1,
		128,
	)

	if p.Size() != 1 {
		t.Fatal("pager failed")
	}
}

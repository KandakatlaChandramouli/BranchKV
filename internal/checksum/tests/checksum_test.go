package tests

import (
	"branchkv-core/internal/checksum"
	"testing"
)

func TestChecksum(
	t *testing.T,
) {

	c := checksum.NewChecksum()

	v := c.Compute(
		[]byte("branchkv"),
	)

	if v == 0 {
		t.Fatal("checksum failed")
	}
}

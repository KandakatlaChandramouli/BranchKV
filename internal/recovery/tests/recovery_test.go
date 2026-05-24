package tests

import (
	"branchkv-core/internal/recovery"
	"testing"
)

func TestRecovery(
	t *testing.T,
) {

	r := recovery.NewRecovery()

	r.Recover()

	if r.Count() != 1 {
		t.Fatal("recovery failed")
	}
}

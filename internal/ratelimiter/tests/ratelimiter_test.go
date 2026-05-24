package tests

import (
	"branchkv-core/internal/ratelimiter"
	"testing"
)

func TestLimiter(
	t *testing.T,
) {

	l := ratelimiter.NewLimiter(2)

	if !l.Allow() {
		t.Fatal("allow failed")
	}
}

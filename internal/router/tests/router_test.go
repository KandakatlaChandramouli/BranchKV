package tests

import (
	"branchkv-core/internal/router"
	"testing"
)

func TestRouter(
	t *testing.T,
) {

	r := router.NewRouter()

	node := r.Route(
		"branch-key",
		4,
	)

	if node >= 4 {
		t.Fatal("invalid route")
	}
}

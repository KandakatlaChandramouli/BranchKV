package tests

import (
	"branchkv-core/internal/embedding"
	"testing"
)

func TestEmbedding(
	t *testing.T,
) {

	e := embedding.NewEmbedding(
		1000,
		8,
	)

	v := e.Encode(1)

	if len(v) != 8 {
		t.Fatal("embedding failed")
	}
}

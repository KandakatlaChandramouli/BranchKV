package tests

import (
	"branchkv-core/internal/vectorindex"
	"testing"
)

func TestVectorIndex(
	t *testing.T,
) {

	idx := vectorindex.NewIndex()

	idx.Insert(
		1,
		[]float32{1, 2, 3},
	)

	if idx.Size() != 1 {
		t.Fatal("vector insert failed")
	}
}

package tests

import (
	"branchkv-core/internal/vectordb"
	"testing"
)

func TestVectorDB(
	t *testing.T,
) {

	db := vectordb.NewVectorDB()

	db.Insert(
		1,
		[]float32{
			1, 2, 3,
		},
	)

	d := db.Distance(
		[]float32{1, 2},
		[]float32{1, 2},
	)

	if d != 0 {
		t.Fatal("distance failed")
	}
}

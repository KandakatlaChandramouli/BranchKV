package tests

import (
	"branchkv-core/internal/objectstore"
	"testing"
)

func TestObjectStore(
	t *testing.T,
) {

	store := objectstore.NewObjectStore()

	store.Put(
		"key",
		[]byte("value"),
	)

	_, ok := store.Get("key")

	if !ok {
		t.Fatal("object missing")
	}
}

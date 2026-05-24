package tests

import (
	"branchkv-core/internal/decoder"
	"branchkv-core/internal/types"
	"testing"
)

func TestDecoder(
	t *testing.T,
) {

	d := decoder.NewDecoder()

	v := d.Decode(
		[]types.TokenID{
			1, 2, 3,
		},
	)

	if v != 4 {
		t.Fatal("decode failed")
	}
}

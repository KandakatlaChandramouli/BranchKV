package simd

import (
	"branchkv-core/internal/types"
)

func CompareTokens(
	a []types.TokenID,
	b []types.TokenID,
) bool {

	if len(a) != len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {

		if a[i] != b[i] {
			return false
		}
	}

	return true
}

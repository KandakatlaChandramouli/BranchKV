package specmerge

import (
	"branchkv-core/internal/virtual_mem/delta"
)

func Merge(
	left []delta.Delta,
	right []delta.Delta,
) []delta.Delta {

	out := make(
		[]delta.Delta,
		0,
		len(left)+len(right),
	)

	out = append(
		out,
		left...,
	)

	out = append(
		out,
		right...,
	)

	return out
}

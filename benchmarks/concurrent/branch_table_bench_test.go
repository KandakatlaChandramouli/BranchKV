package concurrent

import (
	"branchkv-core/internal/concurrent"
	"branchkv-core/internal/tree"
	"strconv"
	"testing"
)

func BenchmarkBranchTable(
	b *testing.B,
) {

	bt := concurrent.NewBranchTable()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		id := uint64(i)

		bt.Store(
			id,
			&tree.Branch{
				ID: id,
			},
		)

		_, _ = bt.Load(id)

		_ = strconv.Itoa(i)
	}
}

package distributed_checkpoint

import (
	"branchkv-core/internal/checkpointing"
	"testing"
)

func BenchmarkCheckpointFabric(
	b *testing.B,
) {

	manager := checkpointing.NewManager()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		manager.Save(
			uint64(i),
		)
	}
}

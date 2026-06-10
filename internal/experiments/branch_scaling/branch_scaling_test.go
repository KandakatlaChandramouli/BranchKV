package branch_scaling

import (
    "testing"
)

func createBranch() {
    // actual BranchKV branch creation call
}

func BenchmarkBranchScaling(b *testing.B) {
    b.ResetTimer()

    for i := 0; i < b.N; i++ {
        createBranch()
    }
}

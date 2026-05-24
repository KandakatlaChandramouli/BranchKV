package distributed

import (
	"branchkv-core/internal/loadbalancer"
	"branchkv-core/internal/router"
	"testing"
)

func BenchmarkDistributedRouting(
	b *testing.B,
) {

	r := router.NewRouter()

	l := loadbalancer.NewLoadBalancer()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		_ = r.Route(
			"branch-key",
			16,
		)

		_ = l.Next(16)
	}
}

package tests

import (
	"branchkv-core/internal/metrics"
	"testing"
)

func TestMetrics(
	t *testing.T,
) {

	m := metrics.NewMetrics()

	m.RecordFork()
	m.RecordWrite()

	if m.Forks.Load() != 1 {
		t.Fatal("fork metric failed")
	}

	if m.Writes.Load() != 1 {
		t.Fatal("write metric failed")
	}
}

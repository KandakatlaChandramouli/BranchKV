package specmetrics

import "sync/atomic"

type Metrics struct {
	faults    atomic.Uint64
	commits   atomic.Uint64
	rollbacks atomic.Uint64
}

func NewMetrics() *Metrics {

	return &Metrics{}
}

func (m *Metrics) Fault() {

	m.faults.Add(1)
}

func (m *Metrics) Commit() {

	m.commits.Add(1)
}

func (m *Metrics) Rollback() {

	m.rollbacks.Add(1)
}

func (m *Metrics) Snapshot() (
	uint64,
	uint64,
	uint64,
) {

	return m.faults.Load(),
		m.commits.Load(),
		m.rollbacks.Load()
}

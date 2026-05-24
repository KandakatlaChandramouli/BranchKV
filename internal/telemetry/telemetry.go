package telemetry

import (
	"sync/atomic"
	"time"
)

type Telemetry struct {
	Events atomic.Uint64
	Start  time.Time
}

func NewTelemetry() *Telemetry {

	return &Telemetry{
		Start: time.Now(),
	}
}

func (t *Telemetry) Record() {
	t.Events.Add(1)
}

func (t *Telemetry) Uptime() time.Duration {
	return time.Since(t.Start)
}

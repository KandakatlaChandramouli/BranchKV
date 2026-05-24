package runtime_profiler

import (
	"sync/atomic"
	"time"
)

type Profiler struct {
	events atomic.Uint64
	start  time.Time
}

func NewProfiler() *Profiler {

	return &Profiler{
		start: time.Now(),
	}
}

func (p *Profiler) Record() {
	p.events.Add(1)
}

func (p *Profiler) Events() uint64 {
	return p.events.Load()
}

func (p *Profiler) Uptime() time.Duration {
	return time.Since(p.start)
}

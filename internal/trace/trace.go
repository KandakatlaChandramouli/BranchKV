package trace

import (
	"sync"
)

type Span struct {
	Name string
}

type Tracer struct {
	spans []Span
	mu    sync.Mutex
}

func NewTracer() *Tracer {

	return &Tracer{
		spans: make([]Span, 0),
	}
}

func (t *Tracer) Start(
	name string,
) {

	t.mu.Lock()
	defer t.mu.Unlock()

	t.spans = append(
		t.spans,
		Span{
			Name: name,
		},
	)
}

func (t *Tracer) Size() int {

	t.mu.Lock()
	defer t.mu.Unlock()

	return len(t.spans)
}

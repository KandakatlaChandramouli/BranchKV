package eventbus

import (
	"sync"
)

type Event struct {
	Name string
}

type Bus struct {
	events []Event
	mu     sync.Mutex
}

func NewBus() *Bus {

	return &Bus{
		events: make([]Event, 0),
	}
}

func (b *Bus) Publish(
	name string,
) {

	b.mu.Lock()
	defer b.mu.Unlock()

	b.events = append(
		b.events,
		Event{
			Name: name,
		},
	)
}

func (b *Bus) Size() int {

	b.mu.Lock()
	defer b.mu.Unlock()

	return len(b.events)
}

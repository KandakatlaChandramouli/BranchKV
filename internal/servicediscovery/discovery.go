package servicediscovery

import (
	"sync"
)

type Service struct {
	Name string
	Host string
}

type Discovery struct {
	services map[string]Service
	mu       sync.RWMutex
}

func NewDiscovery() *Discovery {

	return &Discovery{
		services: make(
			map[string]Service,
		),
	}
}

func (d *Discovery) Register(
	s Service,
) {

	d.mu.Lock()
	defer d.mu.Unlock()

	d.services[s.Name] = s
}

func (d *Discovery) Resolve(
	name string,
) (Service, bool) {

	d.mu.RLock()
	defer d.mu.RUnlock()

	v, ok := d.services[name]

	return v, ok
}

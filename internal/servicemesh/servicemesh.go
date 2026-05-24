package servicemesh

import (
	"sync"
)

type Route struct {
	Source string
	Target string
}

type Mesh struct {
	routes []Route
	mu     sync.Mutex
}

func NewMesh() *Mesh {

	return &Mesh{
		routes: make([]Route, 0),
	}
}

func (m *Mesh) Connect(
	src string,
	dst string,
) {

	m.mu.Lock()
	defer m.mu.Unlock()

	m.routes = append(
		m.routes,
		Route{
			Source: src,
			Target: dst,
		},
	)
}

func (m *Mesh) Size() int {

	m.mu.Lock()
	defer m.mu.Unlock()

	return len(m.routes)
}

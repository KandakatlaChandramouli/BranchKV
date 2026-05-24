package manifest

import (
	"sync"
)

type File struct {
	Name string
}

type Manifest struct {
	files []File
	mu    sync.Mutex
}

func NewManifest() *Manifest {

	return &Manifest{
		files: make([]File, 0),
	}
}

func (m *Manifest) Add(
	name string,
) {

	m.mu.Lock()
	defer m.mu.Unlock()

	m.files = append(
		m.files,
		File{
			Name: name,
		},
	)
}

func (m *Manifest) Size() int {

	m.mu.Lock()
	defer m.mu.Unlock()

	return len(m.files)
}

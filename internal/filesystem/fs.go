package filesystem

import (
	"sync"
)

type File struct {
	Name string
	Data []byte
}

type FileSystem struct {
	files map[string]File
	mu    sync.RWMutex
}

func NewFileSystem() *FileSystem {

	return &FileSystem{
		files: make(
			map[string]File,
		),
	}
}

func (f *FileSystem) Create(
	name string,
	data []byte,
) {

	f.mu.Lock()
	defer f.mu.Unlock()

	f.files[name] = File{
		Name: name,
		Data: data,
	}
}

func (f *FileSystem) Read(
	name string,
) ([]byte, bool) {

	f.mu.RLock()
	defer f.mu.RUnlock()

	v, ok := f.files[name]

	return v.Data, ok
}

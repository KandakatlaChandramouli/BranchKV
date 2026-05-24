package journal

import (
	"os"
	"sync"
)

type Journal struct {
	file *os.File
	mu   sync.Mutex
}

func NewJournal(
	path string,
) (*Journal, error) {

	f, err := os.Create(path)

	if err != nil {
		return nil, err
	}

	return &Journal{
		file: f,
	}, nil
}

func (j *Journal) Write(
	data []byte,
) error {

	j.mu.Lock()
	defer j.mu.Unlock()

	_, err := j.file.Write(data)

	return err
}

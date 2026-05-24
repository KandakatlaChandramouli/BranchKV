package wal

import (
	"encoding/binary"
	"os"
	"sync"
)

type WAL struct {
	file *os.File
	mu   sync.Mutex
}

func NewWAL(
	path string,
) (*WAL, error) {

	f, err := os.OpenFile(
		path,
		os.O_CREATE|
			os.O_RDWR|
			os.O_APPEND,
		0644,
	)

	if err != nil {
		return nil, err
	}

	return &WAL{
		file: f,
	}, nil
}

func (w *WAL) Append(
	value uint64,
) error {

	w.mu.Lock()
	defer w.mu.Unlock()

	buf := make([]byte, 8)

	binary.LittleEndian.PutUint64(
		buf,
		value,
	)

	_, err := w.file.Write(buf)

	return err
}

func (w *WAL) Close() error {
	return w.file.Close()
}

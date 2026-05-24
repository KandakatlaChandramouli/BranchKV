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

func Open(
	path string,
) (
	*WAL,
	error,
) {

	file, err := os.OpenFile(
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
		file: file,
	}, nil
}

func NewWAL(
	path string,
) (
	*WAL,
	error,
) {

	return Open(path)
}

func (w *WAL) Append(
	data []byte,
) error {

	w.mu.Lock()
	defer w.mu.Unlock()

	size := uint64(len(data))

	header := make([]byte, 8)

	binary.LittleEndian.PutUint64(
		header,
		size,
	)

	_, err := w.file.Write(
		header,
	)

	if err != nil {
		return err
	}

	_, err = w.file.Write(
		data,
	)

	if err != nil {
		return err
	}

	return w.file.Sync()
}

func (w *WAL) Close() error {
	return w.file.Close()
}

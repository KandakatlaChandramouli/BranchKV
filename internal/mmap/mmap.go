package mmap

import (
	"os"
	"syscall"
)

type MMapRegion struct {
	File *os.File
	Data []byte
	Size int
}

func NewMMapRegion(
	path string,
	size int,
) (*MMapRegion, error) {

	file, err := os.Create(path)
	if err != nil {
		return nil, err
	}

	err = file.Truncate(int64(size))
	if err != nil {
		return nil, err
	}

	data, err := syscall.Mmap(
		int(file.Fd()),
		0,
		size,
		syscall.PROT_READ|
			syscall.PROT_WRITE,
		syscall.MAP_SHARED,
	)

	if err != nil {
		return nil, err
	}

	return &MMapRegion{
		File: file,
		Data: data,
		Size: size,
	}, nil
}

func (m *MMapRegion) Close() error {

	err := syscall.Munmap(m.Data)
	if err != nil {
		return err
	}

	return m.File.Close()
}

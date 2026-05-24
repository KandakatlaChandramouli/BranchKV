package mmapengine

import (
	"os"
	"syscall"
)

type Region struct {
	Data []byte
}

func MapFile(
	path string,
	size int,
) (
	*Region,
	error,
) {

	file, err := os.OpenFile(
		path,
		os.O_CREATE|
			os.O_RDWR,
		0644,
	)

	if err != nil {
		return nil, err
	}

	err = file.Truncate(
		int64(size),
	)

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

	return &Region{
		Data: data,
	}, nil
}

func (r *Region) Unmap() error {

	return syscall.Munmap(
		r.Data,
	)
}

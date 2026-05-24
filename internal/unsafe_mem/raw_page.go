package unsafe_mem

import (
	"unsafe"
)

type RawPage struct {
	Ptr  unsafe.Pointer
	Size uintptr
}

func NewRawPage(size uintptr) *RawPage {

	mem := make([]byte, size)

	return &RawPage{
		Ptr:  unsafe.Pointer(&mem[0]),
		Size: size,
	}
}

func (r *RawPage) WriteFloat32(
	offset uintptr,
	value float32,
) {

	ptr := (*float32)(
		unsafe.Pointer(
			uintptr(r.Ptr) + offset,
		),
	)

	*ptr = value
}

func (r *RawPage) ReadFloat32(
	offset uintptr,
) float32 {

	ptr := (*float32)(
		unsafe.Pointer(
			uintptr(r.Ptr) + offset,
		),
	)

	return *ptr
}

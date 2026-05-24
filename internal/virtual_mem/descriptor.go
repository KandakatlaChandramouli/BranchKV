package virtual_mem

func NewDescriptor(
	virtual uint64,
	physical uint64,
	offset uint64,
	length uint64,
) VirtualDescriptor {

	return VirtualDescriptor{
		VirtualPageID:  virtual,
		PhysicalPageID: physical,
		Offset:         offset,
		Length:         length,
	}
}

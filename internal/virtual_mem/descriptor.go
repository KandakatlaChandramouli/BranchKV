package virtual_mem

type VirtualDescriptor struct {
	LogicalID uint64
	Page      *PhysicalPage
}

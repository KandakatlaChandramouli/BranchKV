package virtual_mem

import "errors"

var ErrPageFault = errors.New(
	"virtual page missing",
)

type PageFaultHandler struct {
	table *PageTable
}

func NewPageFaultHandler(
	table *PageTable,
) *PageFaultHandler {

	return &PageFaultHandler{
		table: table,
	}
}

func (p *PageFaultHandler) Resolve(
	virtual uint64,
) (
	*PhysicalFrame,
	error,
) {

	frame, ok := p.table.Resolve(
		virtual,
	)

	if !ok {
		return nil, ErrPageFault
	}

	return frame, nil
}

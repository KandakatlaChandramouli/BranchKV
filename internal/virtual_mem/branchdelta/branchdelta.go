package branchdelta

import "branchkv-core/internal/virtual_mem/delta"

type Runtime struct {
	entries []delta.Delta
}

func NewRuntime() *Runtime {

	return &Runtime{}
}

func (r *Runtime) Add(
	d delta.Delta,
) {

	r.entries = append(
		r.entries,
		d,
	)
}

func (r *Runtime) Size() int {

	return len(r.entries)
}

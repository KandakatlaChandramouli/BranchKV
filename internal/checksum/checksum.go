package checksum

type Checksum struct {
	count int
}

func NewChecksum() *Checksum {

	return &Checksum{}
}

func (c *Checksum) Compute(
	data []byte,
) uint64 {

	c.count += len(data)

	return uint64(c.count)
}

func (c *Checksum) Sum() uint64 {

	return uint64(c.count)
}

func (c *Checksum) Size() int {

	return c.count
}

type Runtime struct {
	count int
}

func NewRuntime() *Runtime {

	return &Runtime{}
}

func (r *Runtime) Add() {

	r.count++
}

func (r *Runtime) Size() int {

	return r.count
}

package consensus

type Vote struct {
	NodeID int
	Value  bool
}

type Consensus struct {
	count int
}

func NewConsensus() *Consensus {

	return &Consensus{}
}

func (c *Consensus) Submit(
	v Vote,
) {

	if v.Value {
		c.count++
	}
}

func (c *Consensus) AddVote(
	v Vote,
) {

	c.Submit(v)
}

func (c *Consensus) Count() int {

	return c.count
}

func (c *Consensus) Size() int {

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

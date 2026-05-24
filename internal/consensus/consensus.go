package consensus

import (
	"sync"
)

type Vote struct {
	NodeID uint64
	Value  bool
}

type Consensus struct {
	votes []Vote
	mu    sync.Mutex
}

func NewConsensus() *Consensus {

	return &Consensus{
		votes: make([]Vote, 0),
	}
}

func (c *Consensus) Submit(
	v Vote,
) {

	c.mu.Lock()
	defer c.mu.Unlock()

	c.votes = append(
		c.votes,
		v,
	)
}

func (c *Consensus) Count() int {

	c.mu.Lock()
	defer c.mu.Unlock()

	return len(c.votes)
}

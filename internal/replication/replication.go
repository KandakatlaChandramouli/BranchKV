package replication

import (
	"branchkv-core/internal/network"
	"sync"
)

type Replicator struct {
	log []network.Message
	mu  sync.Mutex
}

func NewReplicator() *Replicator {

	return &Replicator{
		log: make(
			[]network.Message,
			0,
		),
	}
}

func (r *Replicator) Replicate(
	msg network.Message,
) {

	r.mu.Lock()
	defer r.mu.Unlock()

	r.log = append(
		r.log,
		msg,
	)
}

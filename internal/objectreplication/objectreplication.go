package objectreplication

import (
	"sync"
)

type Replica struct {
	ID uint64
}

type Replication struct {
	replicas []Replica
	mu       sync.Mutex
}

func NewReplication() *Replication {

	return &Replication{
		replicas: make([]Replica, 0),
	}
}

func (r *Replication) Add(
	id uint64,
) {

	r.mu.Lock()
	defer r.mu.Unlock()

	r.replicas = append(
		r.replicas,
		Replica{
			ID: id,
		},
	)
}

func (r *Replication) Size() int {

	r.mu.Lock()
	defer r.mu.Unlock()

	return len(r.replicas)
}

package replication

import "branchkv-core/internal/network"

type Replicator struct {
	count int
}

func NewReplicator() *Replicator {

	return &Replicator{}
}

func (r *Replicator) Replicate(
	_ network.Message,
) {
	r.count++
}

func (r *Replicator) Size() int {
	return r.count
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

package distributedwal

type Replica struct {
	ID uint64
}

type DistributedWAL struct {
	replicas []Replica
}

func NewDistributedWAL() *DistributedWAL {

	return &DistributedWAL{
		replicas: make(
			[]Replica,
			0,
		),
	}
}

func (d *DistributedWAL) AddReplica(
	id uint64,
) {

	d.replicas = append(
		d.replicas,
		Replica{
			ID: id,
		},
	)
}

func (d *DistributedWAL) Replicas() int {
	return len(d.replicas)
}

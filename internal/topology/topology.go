package topology

type Link struct {
	From uint64
	To   uint64
}

type Topology struct {
	links []Link
}

func NewTopology() *Topology {

	return &Topology{
		links: make([]Link, 0),
	}
}

func (t *Topology) Connect(
	from uint64,
	to uint64,
) {

	t.links = append(
		t.links,
		Link{
			From: from,
			To:   to,
		},
	)
}

func (t *Topology) Size() int {
	return len(t.links)
}

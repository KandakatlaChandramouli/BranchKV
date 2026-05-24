package nvlink

type Link struct {
	GPUA int
	GPUB int
}

type NVLinkFabric struct {
	links []Link
}

func NewNVLinkFabric() *NVLinkFabric {

	return &NVLinkFabric{
		links: make([]Link, 0),
	}
}

func (n *NVLinkFabric) Connect(
	a int,
	b int,
) {

	n.links = append(
		n.links,
		Link{
			GPUA: a,
			GPUB: b,
		},
	)
}

func (n *NVLinkFabric) Size() int {
	return len(n.links)
}

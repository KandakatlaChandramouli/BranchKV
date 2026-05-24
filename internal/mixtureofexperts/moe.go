package mixtureofexperts

type Expert struct {
	ID uint64
}

type MoE struct {
	experts []Expert
}

func NewMoE() *MoE {

	return &MoE{
		experts: make([]Expert, 0),
	}
}

func (m *MoE) AddExpert(
	id uint64,
) {

	m.experts = append(
		m.experts,
		Expert{
			ID: id,
		},
	)
}

func (m *MoE) Size() int {
	return len(m.experts)
}

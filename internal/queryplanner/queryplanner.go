package queryplanner

type Plan struct {
	Operations []string
}

func NewPlan() *Plan {

	return &Plan{
		Operations: make(
			[]string,
			0,
		),
	}
}

func (p *Plan) Add(
	op string,
) {

	p.Operations = append(
		p.Operations,
		op,
	)
}

func (p *Plan) Steps() int {
	return len(p.Operations)
}

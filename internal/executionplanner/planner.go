package executionplanner

type Step struct {
	Name string
}

type Planner struct {
	steps []Step
}

func NewPlanner() *Planner {

	return &Planner{
		steps: make([]Step, 0),
	}
}

func (p *Planner) Add(
	name string,
) {

	p.steps = append(
		p.steps,
		Step{
			Name: name,
		},
	)
}

func (p *Planner) Size() int {
	return len(p.steps)
}

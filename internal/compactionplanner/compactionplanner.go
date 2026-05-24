package compactionplanner

type Plan struct {
	Level int
}

type Planner struct {
	plans []Plan
}

func NewPlanner() *Planner {

	return &Planner{
		plans: make([]Plan, 0),
	}
}

func (p *Planner) AddPlan(
	level int,
) {

	p.plans = append(
		p.plans,
		Plan{
			Level: level,
		},
	)
}

func (p *Planner) Size() int {
	return len(p.plans)
}

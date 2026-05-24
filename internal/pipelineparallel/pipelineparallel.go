package pipelineparallel

type Stage struct {
	ID uint64
}

type Pipeline struct {
	stages []Stage
}

func NewPipeline() *Pipeline {

	return &Pipeline{
		stages: make([]Stage, 0),
	}
}

func (p *Pipeline) AddStage(
	id uint64,
) {

	p.stages = append(
		p.stages,
		Stage{
			ID: id,
		},
	)
}

func (p *Pipeline) Size() int {
	return len(p.stages)
}

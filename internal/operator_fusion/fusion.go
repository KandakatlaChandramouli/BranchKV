package operator_fusion

type Operator struct {
	Name string
}

type FusionEngine struct {
	ops []Operator
}

func NewFusionEngine() *FusionEngine {

	return &FusionEngine{
		ops: make([]Operator, 0),
	}
}

func (f *FusionEngine) Fuse(
	name string,
) {

	f.ops = append(
		f.ops,
		Operator{
			Name: name,
		},
	)
}

func (f *FusionEngine) Size() int {
	return len(f.ops)
}

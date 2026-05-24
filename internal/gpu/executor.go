package gpu

type GPUExecutor struct {
	DeviceID int
}

func NewGPUExecutor(
	id int,
) *GPUExecutor {

	return &GPUExecutor{
		DeviceID: id,
	}
}

func (g *GPUExecutor) Execute(
	input []float32,
) []float32 {

	out := make(
		[]float32,
		len(input),
	)

	copy(out, input)

	return out
}

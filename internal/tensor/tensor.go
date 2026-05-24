package tensor

type Tensor struct {
	Shape []int
	Data  []float32
}

func NewTensor(
	shape []int,
	data []float32,
) *Tensor {

	return &Tensor{
		Shape: shape,
		Data:  data,
	}
}

func (t *Tensor) Size() int {
	return len(t.Data)
}

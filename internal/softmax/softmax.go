package softmax

import (
	"math"
)

func Compute(
	input []float64,
) []float64 {

	out := make(
		[]float64,
		len(input),
	)

	var sum float64

	for _, v := range input {
		sum += math.Exp(v)
	}

	for i, v := range input {

		out[i] = math.Exp(v) / sum
	}

	return out
}

package flashattention

import (
	"math"
)

type FlashAttention struct{}

func NewFlashAttention() *FlashAttention {
	return &FlashAttention{}
}

func (f *FlashAttention) Compute(
	q []float64,
	k []float64,
	v []float64,
) []float64 {

	out := make(
		[]float64,
		len(v),
	)

	var score float64

	for i := 0; i < len(q); i++ {
		score += q[i] * k[i]
	}

	score /= math.Sqrt(
		float64(len(q)),
	)

	for i := 0; i < len(v); i++ {
		out[i] = score * v[i]
	}

	return out
}

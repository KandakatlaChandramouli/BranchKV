package rope

import (
	"math"
)

type RoPE struct{}

func NewRoPE() *RoPE {
	return &RoPE{}
}

func (r *RoPE) Rotate(
	value float64,
	pos float64,
) float64 {

	return value *
		math.Cos(pos)
}

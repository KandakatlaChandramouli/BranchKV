package layernorm

func Normalize(
	input []float32,
) []float32 {

	out := make(
		[]float32,
		len(input),
	)

	var mean float32

	for _, v := range input {
		mean += v
	}

	mean /= float32(len(input))

	for i, v := range input {
		out[i] = v - mean
	}

	return out
}

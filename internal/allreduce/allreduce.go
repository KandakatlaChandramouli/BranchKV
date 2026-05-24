package allreduce

func Sum(
	input [][]float32,
) []float32 {

	if len(input) == 0 {
		return nil
	}

	out := make(
		[]float32,
		len(input[0]),
	)

	for _, vec := range input {

		for i, v := range vec {
			out[i] += v
		}
	}

	return out
}

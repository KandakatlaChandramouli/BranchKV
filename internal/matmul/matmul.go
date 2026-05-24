package matmul

func Multiply(
	a [][]float32,
	b [][]float32,
) [][]float32 {

	rows := len(a)

	cols := len(b[0])

	k := len(b)

	out := make(
		[][]float32,
		rows,
	)

	for i := 0; i < rows; i++ {

		out[i] = make(
			[]float32,
			cols,
		)

		for j := 0; j < cols; j++ {

			var sum float32

			for x := 0; x < k; x++ {
				sum += a[i][x] * b[x][j]
			}

			out[i][j] = sum
		}
	}

	return out
}

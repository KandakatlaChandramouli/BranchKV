package activation

import (
	"math"
)

func GELU(
	x float64,
) float64 {

	return 0.5 *
		x *
		(1 +
			math.Tanh(
				math.Sqrt(
					2/math.Pi,
				)*
					(x+
						0.044715*
							math.Pow(x, 3)),
			))
}

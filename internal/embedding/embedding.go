package embedding

type Embedding struct {
	VocabSize int
	Dimension int
}

func NewEmbedding(
	vocab int,
	dim int,
) *Embedding {

	return &Embedding{
		VocabSize: vocab,
		Dimension: dim,
	}
}

func (e *Embedding) Encode(
	token uint64,
) []float32 {

	out := make(
		[]float32,
		e.Dimension,
	)

	for i := 0; i < e.Dimension; i++ {
		out[i] = float32(token)
	}

	return out
}

package decoder

import (
	"branchkv-core/internal/beamsearch"
	"branchkv-core/internal/decoder"
	"branchkv-core/internal/types"
	"testing"
)

func BenchmarkDecoderRuntime(
	b *testing.B,
) {

	d := decoder.NewDecoder()

	beam := beamsearch.NewBeamSearch(4)

	seq := []types.TokenID{
		1, 2, 3, 4,
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		_ = d.Decode(seq)

		_ = beam.Expand(seq)
	}
}

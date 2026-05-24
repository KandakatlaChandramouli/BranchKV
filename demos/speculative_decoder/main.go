package main

import (
	"branchkv-core/internal/attentioncache"
	"branchkv-core/internal/beamsearch"
	"branchkv-core/internal/decoder"
	"branchkv-core/internal/pagedattention"
	"branchkv-core/internal/runtime_profiler"
	"branchkv-core/internal/tokenwindow"
	"branchkv-core/internal/trace"
	"branchkv-core/internal/types"
	"fmt"
)

func main() {

	decoderEngine := decoder.NewDecoder()

	beam := beamsearch.NewBeamSearch(4)

	window := tokenwindow.NewWindow(4)

	paged := pagedattention.NewPagedAttention()

	cache := attentioncache.NewCache()

	profiler := runtime_profiler.NewProfiler()

	tracer := trace.NewTracer()

	seq := []types.TokenID{
		1, 2, 3, 4, 5, 6,
	}

	sliced := window.Slice(seq)

	next := decoderEngine.Decode(
		sliced,
	)

	beams := beam.Expand(
		sliced,
	)

	paged.Allocate(1)

	cache.Store(
		1,
		[]float32{
			1, 2, 3,
		},
	)

	profiler.Record()

	tracer.Start("decode")

	fmt.Println(
		"next:",
		next,
	)

	fmt.Println(
		"beams:",
		len(beams),
	)

	fmt.Println(
		"paged:",
		paged.Size(),
	)

	fmt.Println(
		"cache:",
		cache.Size(),
	)

	fmt.Println(
		"events:",
		profiler.Events(),
	)

	fmt.Println(
		"trace:",
		tracer.Size(),
	)
}
